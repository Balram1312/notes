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

--------------------------------------------------
import 'dart:convert';
import 'dart:async';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

class KpiDashboard extends StatefulWidget {
  const KpiDashboard({super.key});

  @override
  State<KpiDashboard> createState() => _KpiDashboardState();
}

class _KpiDashboardState extends State<KpiDashboard> {
  Future<List<dynamic>> kpi3() async {
    var kpi = Uri.parse("http://172.24.0.1:8080/employees/kpi3");
    var res = await http.get(kpi);

    if (res.statusCode == 200) {
      // Request successful
      var jsonData = jsonDecode(res.body);
      debugPrint('data is $jsonData');
      return jsonData;
    } else {
      // Request failed
      throw Exception('Unxexpected error occured');
    }
  }

  Future<List<dynamic>> kpi4() async {
    var kpi = Uri.parse("http://172.24.0.1:8686/employees/kpi4");
    var res = await http.get(kpi);

    if (res.statusCode == 200) {
      // Request successful
      var jsonData = jsonDecode(res.body);
      return jsonData;
    } else {
      // Request failed
      throw Exception('Failed to fetch data');
    }
  }

  Future<List<dynamic>> kpi6() async {
    var kpi = Uri.parse("http://172.24.0.1:8686/employees/kpi6");
    var res = await http.get(kpi);

    if (res.statusCode == 200) {
      // Request successful
      var jsonData = jsonDecode(res.body);
      return jsonData;
    } else {
      // Request failed
      throw Exception('Failed to fetch data');
    }
  }

  @override
  Widget build(BuildContext context) {
    double screenWidth = MediaQuery.of(context).size.width;
    double screenHeight = MediaQuery.of(context).size.height;
    return Scaffold(
      backgroundColor: Colors.teal,
      appBar: AppBar(
        centerTitle: true,
        backgroundColor: Colors.teal,
        elevation: 0,
        leading: IconButton(
            onPressed: () {
              Navigator.pop(context);
            },
            icon: const Icon(Icons.arrow_left)),
        title: const Text(
          'KPI DASHBOARD',
          style: TextStyle(fontSize: 20, letterSpacing: 1.5),
        ),
      ),
      body: SingleChildScrollView(
        child: Center(
          child: Column(
            children: [
              Container(
                decoration: BoxDecoration(
                  borderRadius: BorderRadius.circular(12),
                  color: const Color(0xFFFFFFFF).withOpacity(0.50),
                ),
                child: FutureBuilder<List<dynamic>>(
                  future: kpi3(),
                  builder: (context, snapshot) {
                    if (snapshot.hasData) {
                      return SingleChildScrollView(
                        child: DataTable(
                          columns: const [
                            DataColumn(label: Text('NAME')),
                            DataColumn(label: Text('MOST LEAVE 2023')),
                          ],
                          rows: snapshot.data?.map((data) {
                                return DataRow(cells: [
                                  DataCell(Text(data?['Name'] ?? '')),
                                  DataCell(Text(
                                      data?['Totalleave2023'].toString() ??
                                          '')),
                                ]);
                              }).toList() ??
                              [],
                        ),
                      );
                    } else if (snapshot.hasError) {
                      debugPrint(' error is ${snapshot.hasError}');
                      return Text('Error: ${snapshot.error}');
                    } else {
                      return const Center(child: CircularProgressIndicator());
                    }
                  },
                ),
              )
            ],
          ),
        ),
      ),
    );
  }
}

