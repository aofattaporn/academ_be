package handlers

import (
	"academ_be/models"
	"academ_be/services"
	"bytes"
	"fmt"
	"net/http"
	"net/smtp"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InviteNewMember(c *gin.Context) {

	projectId := c.Param("projectId")
	if projectId == "" {
		handleBussinessError(c, "Can't find your Project ID")
		return
	}

	var inviteReq models.InviteReq
	if err := c.BindJSON(&inviteReq); err != nil {
		handleBussinessError(c, err.Error())
		return
	}

	invitationToken := generateInvitationToken()

	// Retrieve the project by ID
	project, err := services.GetProjectById(c, projectId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	// Send invitation email
	err = sendInvite(inviteReq.InviteEmail, project.ProjectProfile.ProjectName, invitationToken)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	inviteId := primitive.NewObjectID()
	var invite = models.Invite{
		InviteId:     inviteId,
		InviteRoleId: inviteReq.InviteRoleId,
		InviteDate:   inviteReq.InviteDate,
		InviteEmail:  inviteReq.InviteEmail,
		Token:        invitationToken,
	}

	err = services.CreateInvitation(c, projectId, invite)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	memberSetting, err := getProjectMembers(c, projectId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, memberSetting)
}

func generateInvitationToken() string {
	return uuid.New().String()
}

func sendInvite(email, projectName, token string) error {

	// Sender data.
	from := "aofattapon321@gmail.com"
	password := "fyownnkgaikekzjk"

	// Receiver email address.
	to := []string{email}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, err := template.ParseFiles("template.html")
	if err != nil {
		return err
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Invitation to our Event\n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Name        string
		ProjectName string
		AcceptLink  string
	}{
		Name:        email,
		ProjectName: projectName,
		AcceptLink:  "http://localhost:5173/join-project/?token=" + token,
	})

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func DeleteInviteMember(c *gin.Context) {

	projectId := c.Param("projectId")
	if projectId == "" {
		handleBussinessError(c, "Can't find your Project ID")
		return
	}

	inviteId := c.Param("inviteId")
	if projectId == "" {
		handleBussinessError(c, "Can't find your Project ID")
		return
	}

	err := services.DeleteInvitation(c, projectId, inviteId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	memberSetting, err := getProjectMembers(c, projectId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, memberSetting)

}

func AcceptInviteMember(c *gin.Context) {

	projectId := c.Param("projectId")
	if projectId == "" {
		handleBussinessError(c, "Can't find your Project ID")
		return
	}

	token := c.Param("token")
	if projectId == "" {
		handleBussinessError(c, "Can't find your Project ID")
		return
	}

	err := services.DeleteInvitation(c, projectId, token)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	memberSetting, err := getProjectMembers(c, projectId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, memberSetting)

}
