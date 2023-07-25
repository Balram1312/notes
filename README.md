# notes


		
		// Open the file variable which we get from Html form (From Upload File Input)
		src, err := file.Open()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error opening the file: %s", err.Error()))
			return
		}
		defer src.Close()

		
		// Create a temporary file to store the uploaded file
		tempFileName := filepath.Join("./temp", file.Filename)
		dst, err := os.Create(tempFileName)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error creating the destination file: %s", err.Error()))
			return
		}
		defer dst.Close()

		
		// Copy the uploaded file to the destination file
		if _, err = io.Copy(dst, src); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error copying the file: %s", err.Error()))
			return
		}
		//Creating cloudinary connection
		cld, _ := cloudinary.NewFromParams("dwukpu6fw", "728347838492291", "NPX1jbdrKdsW_5L_IZ2iZ6JXKr0")
		
		//Uploading the temporary file to cloudinary cloud
		resp, err := cld.Upload.Upload(c, tempFileName, uploader.UploadParams{PublicID: file.Filename})
		if err != nil {
			log.Println("Failed to upload file to Cloudinary:", err)
			return
		}
		
		//initializing the filename variable with the uploaded file url
		filename = resp.URL


  -----------------------------------------------
  query
  -----------------------------------------------
  SELECT teamname, leavetype, count(*) as leave_count
FROM employee_leave_data
WHERE EXTRACT(YEAR from fromdate) = 2022
AND teamname IN (
    SELECT teamname
    FROM employee_leave_data
    WHERE EXTRACT(YEAR from fromdate) = 2022
    GROUP BY teamname
    ORDER BY count(*) DESC
    LIMIT 2
)
GROUP BY teamname, leavetype;
