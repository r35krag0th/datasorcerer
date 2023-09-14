package routehandlers

import (
	"errors"
	"net/http"

	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/middleware"
	"github.com/danielgtaylor/huma/responses"
	"github.com/r35krag0th/datasorcerer/apimodels"
	"github.com/r35krag0th/datasorcerer/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func RealmRoutes(resource *huma.Resource, db *gorm.DB) {
	root := resource.SubResource("/realms")
	specific := root.SubResource("/{id}")

	root.Get(
		"list-realms",
		"Lists the available realms",
		responses.OK().Model(apimodels.ListRealmsResponse{}),
		responses.InternalServerError(),
	).Run(func(ctx huma.Context) {
		localLogger := middleware.GetLogger(ctx).With("list-realm")
		localLogger.Info(">>> listing realms")
		var output []models.Realm

		result := db.Find(&output)
		if result.Error != nil {
			localLogger.Error("failed to lists realms", zap.Error(result.Error))
			ctx.WriteError(http.StatusInternalServerError, "", result.Error)
			return
		}

		ctx.WriteModel(http.StatusOK, apimodels.ListRealmsResponse{
			Realms: output,
			CountOfObjects: apimodels.CountOfObjects{
				Count: int(result.RowsAffected),
			},
		})
	})

	root.Post(
		"create-realm",
		"Creates a Realm",
		responses.Created().Model(apimodels.CreateRealmResponse{}),
		responses.BadRequest(),
		responses.InternalServerError(),
		responses.Conflict(),
	).Run(func(ctx huma.Context, input apimodels.CreateRealmRequest) {
		localLogger := middleware.GetLogger(ctx).With(
			zap.String("action", "create-realm"),
			zap.String("realm.name", input.Body.Name),
		)
		var output models.Realm

		localLogger.Debug(
			"checking for existing Realms",
		)

		result := db.Model(&models.Realm{}).Find(&output, "name = ?", input.Body.Name)
		if result.RowsAffected > 0 {
			ctx.WriteError(http.StatusConflict, "realm with that name already exists")
			return
		}

		localLogger.Debug(
			"creating Realm",
		)

		output = models.Realm{
			Name: input.Body.Name,
		}

		result = db.Create(&output)
		if result.RowsAffected == 0 || result.Error != nil {
			ctx.WriteError(http.StatusInternalServerError, "failed to create ruleset", result.Error)
			return
		}

		ctx.WriteModel(
			http.StatusCreated,
			apimodels.CreateRealmResponse{
				Realm: output,
			},
		)
	})

	specific.Get(
		"show-realm",
		"",
		responses.OK().Model(apimodels.ShowRealmResponse{}),
		responses.NotFound(),
		responses.InternalServerError(),
	).Run(func(ctx huma.Context, input apimodels.ShowRealmRequest) {
		localLogger := middleware.GetLogger(ctx).With(
			zap.String("action", "show-realm"),
		)
		localLogger.Debug(
			"looking up realm",
			zap.String("realm_id", input.RealmID),
		)
		var output models.Realm
		result := db.Model(&models.Realm{}).
			Find(&output, "id = ?", input.RealmID)

		localLogger.Debug(
			"found realms",
			zap.Int64("count", result.RowsAffected),
			zap.Error(result.Error),
		)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
			ctx.WriteError(http.StatusNotFound, "ruleset was not found")
			return
		} else if result.Error != nil {
			ctx.WriteError(http.StatusInternalServerError, "unexpected issue occurred", result.Error)
			return
		}

		ctx.WriteModel(http.StatusOK, apimodels.ShowRealmResponse{
			Realm: output,
		})
	})

	specific.Delete(
		"delete-realm",
		"Deletes the requested Realm",
		responses.Accepted().Model(apimodels.DeleteRealmResponse{}),
		responses.NotFound(),
		responses.InternalServerError(),
	).Run(func(ctx huma.Context, input apimodels.DeleteRealmRequest) {
		localLogger := middleware.GetLogger(ctx).With(
			zap.String("action", "delete-realm"),
			zap.String("realm.name", input.SpecificRealm.RealmID),
		)
		localLogger.Debug("Looking up the realm")

		var output models.Realm
		result := db.Model(&models.Realm{}).Find(&output, "id = ?", input.RealmID)

		localLogger.Debug("lookup resulted in the following",
			zap.Int64("count", result.RowsAffected),
			zap.Error(result.Error),
		)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
			ctx.WriteError(http.StatusNotFound, "realm was not found")
			return
		} else if result.Error != nil {
			ctx.WriteError(http.StatusInternalServerError, "unexpected issue occurred", result.Error)
			return
		}

		result = db.Model(&output).Delete(&output)
		if result.Error != nil {
			ctx.WriteError(
				http.StatusInternalServerError,
				"unexpected issue occurred",
				result.Error,
			)

			return
		}

		ctx.WriteModel(
			http.StatusAccepted,
			apimodels.DeleteRealmResponse{},
		)
	})

	specific.Put(
		"update-realm",
		"Updates the requested Realm",
		responses.Accepted().Model(apimodels.UpdateRealmResponse{}),
		responses.NotFound(),
		responses.InternalServerError(),
	).Run(func(ctx huma.Context, input apimodels.UpdateRealmRequest) {
		localLogger := middleware.GetLogger(ctx).With(
			zap.String("action", "update-realm"),
			zap.String("realm.name", input.SpecificRealm.RealmID),
		)
		localLogger.Debug("Looking up the realm")

		var output models.Realm
		result := db.Model(&models.Realm{}).Find(&output, "id = ?", input.RealmID)
		if result.RowsAffected == 0 {
			ctx.WriteError(http.StatusNotFound, "realm does not exist")
			return
		}

		if output.Name == input.Body.Name {
			ctx.WriteError(http.StatusNotModified, "the realm had no changes")
			return
		}

		output.Name = input.Body.Name
		result = db.Save(&output)
		if result.Error != nil {
			ctx.WriteError(http.StatusInternalServerError, "failed to save ruleset", result.Error)
			return
		}

		ctx.WriteModel(http.StatusAccepted, apimodels.UpdateRealmResponse{
			Realm: output,
		})
	})
}
