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
)

func getProjectMembers(c *gin.Context, projectId string) (*models.AllMemberProject, error) {
	project, err := services.GetProjectById(c, projectId)
	if err != nil {
		return nil, err
	}

	return &models.AllMemberProject{
		Members: project.Members,
		Roles:   project.Roles,
		Invite:  []models.Invite{},
	}, nil
}

func GetProjectMembers(c *gin.Context) {
	projectId := c.Param("projectsId")
	if projectId == "" {
		handleBussinessError(c, "Can't find your Project ID")
		return
	}

	memberSetting, err := getProjectMembers(c, projectId)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, memberSetting)
}

func ChangeRoleMember(c *gin.Context) {
	projectId := c.Param("projectsId")
	if projectId == "" {
		handleBussinessError(c, "Can't find your Project ID")
		return
	}

	memberId := c.Param("memberId")
	if memberId == "" {
		handleBussinessError(c, "Can't find your Member ID")
		return
	}

	roleId := c.Param("roleId")
	if roleId == "" {
		handleBussinessError(c, "Can't find your Role ID")
		return
	}

	err := services.UpdateRoleByMemberID(c, projectId, memberId, roleId)
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

func InviteNewMember(c *gin.Context) {

	projectId := c.Param("projectsId")
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

	// Send invitation email
	err := sendInvite(inviteReq.InviteEmail, "ProjectName", inviteReq.InviteBy, invitationToken)
	if err != nil {
		handleTechnicalError(c, err.Error())
		return
	}

	var invite = models.Invite{
		InviteRole:  inviteReq.InviteRoleId,
		InviteDate:  inviteReq.InviteDate,
		InviteEmail: inviteReq.InviteEmail,
		Token:       invitationToken,
	}

	handleSuccess(c, http.StatusOK, SUCCESS, GET_MY_PROJECT_SUCCESS, invite)
}

func generateInvitationToken() string {
	return uuid.New().String()
}

func sendInvite(email, projectName, inviteBy, token string) error {

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
		InviteBy    string
		AcceptLink  string
	}{
		Name:        email,
		ProjectName: projectName,
		InviteBy:    inviteBy,
		AcceptLink:  "http://localhost:5173/join-project/?token=" + token,
	})

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		return err
	}

	return nil
}
