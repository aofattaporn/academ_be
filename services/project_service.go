package services

import (
	"academ_be/configs"
	"academ_be/models"
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateProject(c *gin.Context, newUser *models.Project) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	_, err = configs.GetCollection(mongoClient, PROJECT_COLLECTION).InsertOne(ctx, newUser)
	if err != nil {
		return err
	}

	return nil
}

func GetProjectById(c *gin.Context, projectId string) (project *models.ProjectInfo, err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	// Convert projectId string to ObjectId
	objID, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return nil, err
	}

	err = configs.GetCollection(mongoClient, PROJECT_COLLECTION).FindOne(ctx, bson.M{"_id": objID}).Decode(&project)
	if err != nil {
		return nil, err
	}

	return project, err
}

func GetProjectByUserId(c *gin.Context, userId string) (projects []models.AllTasksMyProject, err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	filter := bson.M{"members": bson.M{"$elemMatch": bson.M{"userId": userId}}}

	cursor, err := configs.GetCollection(mongoClient, PROJECT_COLLECTION).Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Decode projects directly into the result slice
	if err := cursor.All(ctx, &projects); err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return []models.AllTasksMyProject{}, nil
	}

	return projects, nil

}

func GetAllProjectCron() (projects []models.ProjectCron, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Define options to specify fields to include
	projection := bson.M{
		"_id":              1,
		"projectProfile":   1,
		"projectStartDate": 1,
		"projectEndDate":   1,
		"isArchive":        1,
		"roles":            1,
		"members":          1,
	}
	opts := options.Find().SetProjection(projection)

	// Find projects matching the filter (empty filter to get all documents)
	cursor, err := configs.GetCollection(mongoClient, PROJECT_COLLECTION).Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, fmt.Errorf("error finding projects: %v", err)
	}
	defer func() {
		if err := cursor.Close(ctx); err != nil {
			fmt.Printf("error closing cursor: %v\n", err)
		}
	}()

	// Decode projects directly into the result slice
	if err := cursor.All(ctx, &projects); err != nil {
		return nil, fmt.Errorf("error decoding projects: %v", err)
	}

	// Optional: Log the projects if needed for debugging
	fmt.Println(projects)
	fmt.Println("==============")

	return projects, nil
}

func GetProjectsByMemberUserID(c *gin.Context, myUserID string) (projects []models.MyProject, err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	// Define filter to find projects where Members.UserID equals myUserID
	filter := bson.M{"members": bson.M{"$elemMatch": bson.M{"userId": myUserID}}}

	// Define options to specify fields to include
	projection := bson.M{"_id": 1, "projectProfile": 1, "members": 1, "projectStartDate": 1, "projectEndDate": 1, "isArchive": 1, "className": 1}
	opts := options.Find().SetProjection(projection)

	// Find projects matching the filter
	cursor, err := configs.GetCollection(mongoClient, PROJECT_COLLECTION).Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Decode projects directly into the result slice
	if err := cursor.All(ctx, &projects); err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return []models.MyProject{}, nil
	}

	return projects, nil
}

func DeleteProjectById(c *gin.Context, projectId string) (project *models.Project, err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	// Convert projectId string to ObjectId
	objID, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return nil, err
	}

	err = configs.GetCollection(mongoClient, PROJECT_COLLECTION).FindOneAndDelete(ctx, bson.M{"_id": objID}).Decode(&project)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func GetProjectDetails(c *gin.Context, projectId string) (projectDetails *models.ProjectDetails, err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	// Convert projectId string to ObjectId
	objID, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return nil, err
	}

	err = configs.GetCollection(mongoClient, PROJECT_COLLECTION).FindOne(ctx, bson.M{"_id": objID}).Decode(&projectDetails)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("project not found")
		}
		return nil, fmt.Errorf("error decoding project: %v", err)
	}

	return projectDetails, nil
}

func UpdateProjectDetails(c *gin.Context, projectId string, projectUpdate models.ProjectDetails) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	// Convert the string task ID to an ObjectID
	id, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "_id", Value: id}}

	update := bson.M{}

	update["className"] = projectUpdate.ClassName
	update["projectProfile"] = projectUpdate.ProjectProfile
	update["views"] = projectUpdate.Views
	// Check if ProjectStartDate is not nil or not set to null
	if projectUpdate.ProjectStartDate != nil {
		update["projectStartDate"] = projectUpdate.ProjectStartDate
	} else {
		update["projectStartDate"] = nil // Set to null in database
	}

	// Check if ProjectEndDate is not nil or not set to null
	if projectUpdate.ProjectEndDate != nil {
		update["projectEndDate"] = projectUpdate.ProjectEndDate
	} else {
		update["projectEndDate"] = nil // Set to null in database
	}

	updateSomeFields := bson.M{"$set": update}

	_, err = configs.GetCollection(mongoClient, PROJECT_COLLECTION).UpdateOne(ctx, filter, updateSomeFields)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProjecArchive(c *gin.Context, projectId string, projectArchive models.ProjectArchive) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	// Convert the string task ID to an ObjectID
	id, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "_id", Value: id}}

	update := bson.M{}

	update["isArchive"] = !projectArchive.IsArchive

	updateSomeFields := bson.M{"$set": update}

	_, err = configs.GetCollection(mongoClient, PROJECT_COLLECTION).UpdateOne(ctx, filter, updateSomeFields)
	if err != nil {
		return err
	}

	return nil
}

func CreateInvitation(c *gin.Context, projectId string, invite models.Invite) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	// Convert projectId to ObjectID
	objID, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return fmt.Errorf("invalid project ID: %v", err)
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$push": bson.M{"invites": invite}}

	// Perform the update on the PROJECT_COLLECTION
	_, err = configs.GetCollection(mongoClient, PROJECT_COLLECTION).UpdateOne(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("project not found")
		}
		return fmt.Errorf("error updating project: %v", err)
	}

	return nil

}

func DeleteInvitation(c *gin.Context, projectId string, inviteId string) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return fmt.Errorf("invalid project ID: %v", err)
	}

	inviteID, err := primitive.ObjectIDFromHex(inviteId)
	if err != nil {
		return fmt.Errorf("invalid project ID: %v", err)
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$pull": bson.M{"invites": bson.M{"inviteId": inviteID}}}

	// Perform the update on the PROJECT_COLLECTION
	_, err = configs.GetCollection(mongoClient, PROJECT_COLLECTION).UpdateOne(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("project not found")
		}
		return fmt.Errorf("error updating project: %v", err)
	}

	return nil

}

func RemoveProjectInviteFromAccept(c *gin.Context, token string) (*primitive.ObjectID, *models.Invite, error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	filter := bson.M{"invites.token": token}
	update := bson.M{"$pull": bson.M{"invites": bson.M{"token": token}}}

	var project models.Project
	err := configs.GetCollection(mongoClient, PROJECT_COLLECTION).FindOneAndUpdate(ctx, filter, update).Decode(&project)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil, fmt.Errorf("project not found")
		}
		return nil, nil, fmt.Errorf("error updating project: %v", err)
	}

	var projectId *primitive.ObjectID = &project.ProjectId
	var invite *models.Invite
	for i, inv := range project.Invites {
		if inv.Token == token {
			invite = &project.Invites[i]
			break
		}
	}

	if invite == nil {
		return nil, nil, fmt.Errorf("invite not found")
	}

	return projectId, invite, nil
}

func AddNewMember(c *gin.Context, projectId primitive.ObjectID, member models.Member) (err error) {

	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": projectId}
	update := bson.M{"$push": bson.M{"members": member}}

	// Perform the update on the PROJECT_COLLECTION
	_, err = configs.GetCollection(mongoClient, PROJECT_COLLECTION).UpdateOne(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("project not found")
		}
		return fmt.Errorf("error updating project: %v", err)
	}

	return nil

}

func RemoveMember(c *gin.Context, projectId string, memberId string) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	// Convert projectId string to ObjectId
	objID, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return fmt.Errorf("invalid project ID: %v", err)
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$pull": bson.M{"members": bson.M{"userId": memberId}}}

	// Perform the update on the PROJECT_COLLECTION
	_, err = configs.GetCollection(mongoClient, PROJECT_COLLECTION).UpdateOne(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("project not found")
		}
		return fmt.Errorf("error updating project: %v", err)
	}

	return nil
}

func UpdateProcessByID(c *gin.Context, projectID string, processID string, process models.Process) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	// Convert projectID and processID to ObjectIDs
	objID, err := primitive.ObjectIDFromHex(projectID)
	if err != nil {
		return fmt.Errorf("invalid project ID: %v", err)
	}

	procID, err := primitive.ObjectIDFromHex(processID)
	if err != nil {
		return fmt.Errorf("invalid process ID: %v", err)
	}

	// Define the filter to match the project and process IDs
	filter := bson.M{"_id": objID, "process.processId": procID}

	// Define the update to set the new process data
	update := bson.M{
		"$set": bson.M{
			"process.$.processName":  process.ProcessName,
			"process.$.processColor": process.ProcessColor,
		},
	}

	// Perform the update on the PROJECT_COLLECTION
	result, err := configs.GetCollection(mongoClient, PROJECT_COLLECTION).UpdateOne(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("project or process not found")
		}
		return fmt.Errorf("error updating process: %v", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("process not found in the project")
	}

	return nil
}

func DeleteProcessbyId(c *gin.Context, projectId string, processId string) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return fmt.Errorf("invalid project ID: %v", err)
	}

	processID, err := primitive.ObjectIDFromHex(processId)
	if err != nil {
		return fmt.Errorf("invalid process ID: %v", err)
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$pull": bson.M{"process": bson.M{"processId": processID}}}

	// Perform the update on the PROJECT_COLLECTION
	_, err = configs.GetCollection(mongoClient, PROJECT_COLLECTION).UpdateOne(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("project not found")
		}
		return fmt.Errorf("error updating project: %v", err)
	}

	return nil

}

func CreateProcessByID(c *gin.Context, projectID string, process models.Process) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	// Convert projectID to ObjectID
	objID, err := primitive.ObjectIDFromHex(projectID)
	if err != nil {
		return fmt.Errorf("invalid project ID: %v", err)
	}

	// Generate a new process ID
	processID := primitive.NewObjectID()

	// Define the filter to match the project
	filter := bson.M{"_id": objID}

	// Define the update to push the new process
	update := bson.M{"$push": bson.M{"process": bson.M{
		"processId":    processID,
		"processName":  process.ProcessName,
		"processColor": process.ProcessColor,
	}}}

	// Perform the update on the PROJECT_COLLECTION
	_, err = configs.GetCollection(mongoClient, PROJECT_COLLECTION).UpdateOne(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("project not found")
		}
		return fmt.Errorf("error creating process: %v", err)
	}

	return nil
}
