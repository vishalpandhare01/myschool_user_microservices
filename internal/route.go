package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vishalpandhare01/internal/handler/admin"
	"github.com/vishalpandhare01/internal/handler/school"
	"github.com/vishalpandhare01/internal/handler/user"
	"github.com/vishalpandhare01/internal/middleware"
)

func Server(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "server running successfully",
	})
}

func RouteSetUp(app *fiber.App) {
	/*-----------------------  user apis -------------------------------*/
	var userRoutes = app.Group("/api/v1/user")
	userRoutes.Post("/register", user.CreateUserHandler) //access for schools and admin
	userRoutes.Post("/sendOtp", user.SendOtp)
	userRoutes.Post("/veryfyOtp", user.VeryfyOtp)

	/*-----------------------  admin apis -------------------------------*/
	var adminRoutes = app.Group("/api/v1/admin", middleware.Authentication, middleware.Authorization)
	adminRoutes.Get("/schools", admin.GetSchoolsHandler)
	adminRoutes.Patch("/makepaidorunpaid/:schoolId", admin.UpdateSchoolByIdHandler)

	/*-----------------------  school apis ------------------------------*/
	var schoolRoutes = app.Group("/api/v1/school", middleware.Authentication)

	schoolRoutes.Get("/", school.GetSchoolByIdHandler)
	//class
	schoolRoutes.Post("/class", school.AddClassHandler)
	schoolRoutes.Get("/class", school.GetClassBySchoolIdHandler)
	schoolRoutes.Delete("/class/:classId", school.DeleteClassByIdHandler)

	//student
	schoolRoutes.Post("/student", school.AddSchoolStudentHandler)
	schoolRoutes.Get("/students", school.GetAllSchoolStudentHandler) // filter with  RegisterNumber , classid , mobile ,email ,fname,lname
	schoolRoutes.Get("/student/:studentId", school.GetStudentByIdHandler)
	schoolRoutes.Put("/student", school.UpdateSchoolStudentHandler)
	schoolRoutes.Delete("/student/:studentId", school.RemoveStudentByIdHandler)
	schoolRoutes.Patch("/student/class/:currentClassId/newclass/:nextClassId", school.MoveBulkStudentToAnotherClassByIdHandler)

	//staff
	schoolRoutes.Post("/staff", school.AddSchoolStaffHandler)
	schoolRoutes.Get("/staff", school.GetAllStaffStudentHandler) // filter with  RegisterNumber , classid , mobile ,email ,fname,lname
	schoolRoutes.Get("/staff/:userid", school.GetStaffbyIdHandler)
	schoolRoutes.Put("/staff", school.UpdateSchoolStaffHandler)
	schoolRoutes.Delete("/staff/:userid", school.DeleteStaffbyIdHandler)

	//fees
	//fee types
	schoolRoutes.Post("/fees/type", school.AddFeesTypeHandler) //Note :- can user for update just need same id
	schoolRoutes.Get("/fees/type", school.GetFeesTypesHandler)
	//fees structure
	schoolRoutes.Post("/fees/structrue", school.AddFeeStructureTypesHandler) //Note :- can user for update just need same id
	schoolRoutes.Get("/fees/structrue", school.GetFeeStructuresHandler)
	//student fees
	schoolRoutes.Post("/students/fees", school.AddStudentFeesHandler)
	schoolRoutes.Get("/students/fees", school.GetStudentFeesHandler)
	//time table
	schoolRoutes.Post("/time-table/", school.CreateTimeTableHandler)
	schoolRoutes.Get("/time-table/:classId", school.GetTimeTableHandler)
	schoolRoutes.Delete("/time-table/:tableId", school.DeleteTimeTableHandler)
}
