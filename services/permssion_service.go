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
)

func CreatePermission(c *gin.Context, permission *models.Permission, permissionId primitive.ObjectID) (err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	permission.Id = permissionId
	_, err = configs.GetCollection(mongoClient, PERMISSION_COLLECTION).InsertOne(ctx, permission)
	if err != nil {
		return err
	}

	return nil
}

func GetPermission(c *gin.Context, permissionId primitive.ObjectID) (permission *models.Permission, err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	err = configs.GetCollection(mongoClient, PERMISSION_COLLECTION).FindOne(ctx, bson.M{"_id": permissionId}).Decode(&permission)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("permission not found")
		}
		return nil, fmt.Errorf("error decoding permission: %v", err)
	}

	return permission, nil
}

func CreateNewRole(c *gin.Context, projectId string, newRole models.Role) error {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	// Convert projectId to ObjectID
	objID, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return fmt.Errorf("invalid project ID: %v", err)
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$push": bson.M{"roles": newRole}}

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

func UpdateRoleName(c *gin.Context, projectId string, roleId string, roleName string) error {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return fmt.Errorf("invalid project ID: %v", err)
	}

	roleObjID, err := primitive.ObjectIDFromHex(roleId)
	if err != nil {
		return fmt.Errorf("invalid role ID: %v", err)
	}

	filter := bson.M{"_id": objID, "roles.roleId": roleObjID}
	update := bson.M{"$set": bson.M{"roles.$.roleName": roleName}}

	_, err = configs.GetCollection(mongoClient, PROJECT_COLLECTION).UpdateOne(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("project or role not found")
		}
		return fmt.Errorf("error updating role name: %v", err)
	}

	return nil
}

func DeleteRole(c *gin.Context, projectId string, roleId string) (project *models.Project, err error) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	// Convert projectId to ObjectID
	objID, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return nil, fmt.Errorf("invalid project ID: %v", err)
	}

	// Convert roleId to ObjectID
	roleObjID, err := primitive.ObjectIDFromHex(roleId)
	if err != nil {
		return nil, fmt.Errorf("invalid role ID: %v", err)
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$pull": bson.M{"roles": bson.M{"roleId": roleObjID}}}

	collection := configs.GetCollection(mongoClient, PROJECT_COLLECTION)
	err = collection.FindOneAndUpdate(ctx, filter, update).Decode(&project)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("project not found")
		}
		return nil, fmt.Errorf("error deleting role: %v", err)
	}

	return project, nil
}

func UpdatePermission(c *gin.Context, permissionId string, updatePermission models.Permission) error {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(permissionId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatePermission}

	_, err = configs.GetCollection(mongoClient, PERMISSION_COLLECTION).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("error updating document: %v", err)
	}

	return nil
}

func UpdateRoleByMemberID(c *gin.Context, projectId string, memberId string, newRoleId string) error {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	projectID, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return err
	}

	newRoleID, err := primitive.ObjectIDFromHex(newRoleId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": projectID, "members.userId": memberId}
	update := bson.M{"$set": bson.M{"members.$.roleId": newRoleID}}

	_, err = configs.GetCollection(mongoClient, PROJECT_COLLECTION).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func DeletePermissionBy(c *gin.Context, permissionId primitive.ObjectID) error {

	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: permissionId}}
	_, err := configs.GetCollection(mongoClient, PERMISSION_COLLECTION).DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil

}
