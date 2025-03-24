package repository

//create attendance by teacher staff

//get attendance by date + school id , subject , teacher id ,student id , classid

//patch to update attendance

/*
 Parse the date string into a time.Time object
		searchDate, err := time.Parse("2006-01-02", dateParam)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid date format, please use YYYY-MM-DD.")
		}

		var items []Item
		err = db.Where("DATE(created_at) = ?", searchDate.Format("2006-01-02")).Find(&items).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error fetching items.")
		}
*/
