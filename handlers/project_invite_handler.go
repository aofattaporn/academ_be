package handlers

import (
	"academ_be/models"
	"academ_be/services"
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/smtp"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func generateInvitationToken() string {
	return uuid.New().String()
}

func sendInvite(email, projectName, token string) error {

	// Sender data.
	from := "academ.projex@gmail.com"
	password := "alhsjlsqtvhyfmal"

	// Receiver email address.
	to := []string{email}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Invitation to our Event\n%s\n\n", mimeHeaders)))

	// TODO : Send Mail Invited Project

	// TODO : Send Mail Tasks Update

	t, err := template.ParseFiles("template.html")
	if err != nil {
		return errors.New("Can't to send this eamil")
	}

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
		return errors.New("Can't to send this eamil")
	}

	return nil
}

func InviteNewMember(c *gin.Context) {

	userID := c.MustGet(USER_ID).(string)
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

	err := checkEmailExisting(c, projectId, inviteReq.InviteEmail)
	if err != nil {
		handleBussinessError(c, "Email Already Existing")
		return
	}

	invitationToken := generateInvitationToken()
	project, err := services.GetProjectById(c, projectId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

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

	membersAndPermission := getMemberAndMemberPermission(c, projectId, userID)

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, membersAndPermission)
}

func DeleteInviteMember(c *gin.Context) {

	userID := c.MustGet(USER_ID).(string)
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

	membersAndPermission := getMemberAndMemberPermission(c, projectId, userID)

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, membersAndPermission)

}

func AcceptInviteMember(c *gin.Context) {

	token := c.Param("token")
	if token == "" {
		handleBussinessError(c, "Can't find your Project ID")
		return
	}

	projectId, invite, err := services.RemoveProjectInviteFromAccept(c, token)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	// Extract the user ID from the request context
	userID := c.MustGet(USER_ID).(string)
	user, err := services.FindUserFullInfoOneById(c, userID)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	inviteRoleID, err := primitive.ObjectIDFromHex(invite.InviteRoleId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	newMember := models.Member{
		UserId:      user.Id,
		UserName:    user.FullName,
		Emaill:      user.Email,
		RoleId:      inviteRoleID,
		AvatarColor: user.AvatarColor,
	}

	err = services.AddNewMember(c, *projectId, newMember)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, "success")

}

func checkEmailExisting(c *gin.Context, projectId string, reqEmail string) error {

	allMemer, err := getProjectMembers(c, projectId)
	if err != nil {
		return err
	}

	for _, m := range allMemer.Members {
		if m.Emaill == reqEmail {
			return errors.New("Eamil Already Existing")
		}
	}

	for _, m := range allMemer.Invites {
		if m.InviteEmail == reqEmail {
			return errors.New("Eamil Already Invite")
		}
	}

	return nil
}
