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




--------------------------------------------
import 'dart:convert';
import 'dart:async';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:fl_chart/fl_chart.dart';

class KpiDashboard extends StatefulWidget {
  const KpiDashboard({super.key});

  @override
  State<KpiDashboard> createState() => _KpiDashboardState();
}

class _KpiDashboardState extends State<KpiDashboard> {
  Future<List<dynamic>> kpi3() async {
    var kpi = Uri.parse("http://172.18.0.1:8080/employees/kpi3");

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
    var kpi = Uri.parse("http://172.18.0.1:8686/employees/kpi4");

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
    var kpi = Uri.parse("http://172.18.0.1:8080/employees/kpi6");

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


   Widget buildBarChart(Map<String, dynamic>? kpi4Data) {
    if (kpi4Data == null) {
      return Container(); // Blank space for the bar chart
    }

    // Your code to build the bar chart using kpi4Data...
    // Please refer to the previous example for building the bar chart.
    // The code inside this function will be specific to your chart implementation.

    // For example, you can create a bar chart using FlChart library as shown before.
    // If you need further assistance with your specific chart implementation,
    // feel free to ask!

    return Container(); // Replace this with your actual bar chart widget
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
              ),
              Container(
                padding: EdgeInsets.all(16),
                decoration: BoxDecoration(
                  borderRadius: BorderRadius.circular(12),
                  color: const Color(0xFFFFFFFF).withOpacity(0.50),
                ),
                child: FutureBuilder<Map<String, dynamic>>(
                  future: kpi4(),
                  builder: (context, snapshot) {
                    if (snapshot.connectionState == ConnectionState.waiting) {
                      return Container(); // Blank space for the bar chart while data is loading
                    } else if (snapshot.hasData) {
                      return buildBarChart(snapshot.data);
                    } else if (snapshot.hasError) {
                      debugPrint('error is ${snapshot.hasError}');
                      return Text('Error: ${snapshot.error}');
                    } else {
                      return Container(); // Blank space for the bar chart if there's no data
                    }
                  },
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
--------------------------------------------------------------
flutter
--------------------------------------------------------------
import 'dart:convert';
import 'dart:async';
import 'package:charts_flutter/flutter.dart' as charts;
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:pie_chart/pie_chart.dart' as piechart;

class Sales {
  final String x;

  final int y;

  Sales(this.x, this.y);
}

class KpiDashboard extends StatefulWidget {
  const KpiDashboard({super.key});

  @override
  State<KpiDashboard> createState() => _KpiDashboardState();
}

class _KpiDashboardState extends State<KpiDashboard> {
  List<charts.Series<dynamic, String>> seriesList = [];
  Map<String, Map<String, double>> allTeamsDataMap = {};

  Map<String, double> aiDataMap = {};
  Map<String, double> itDataMap = {};


Map<String, Map<String, double>> dataMap(List<dynamic> kpi6Data) {
  Map<String, Map<String, double>> teamDataMap = {};
  

  for (var entry in kpi6Data) {
    String teamName = entry['Teamname'];
    String leaveType = entry['Leavetype'];
    double leaveCount = entry['Leavecount'].toDouble();

    if (!teamDataMap.containsKey(teamName)) {
      teamDataMap[teamName] = {};
    }
    teamDataMap[teamName]![leaveType] = leaveCount;
  }

  return teamDataMap;
}

  Future<List<charts.Series<Sales, String>>> _createRandomData() async {
    var values = await kpi4();

    List<Sales> desktopSalesData = [];

    for (var value in values) {
      desktopSalesData.add(Sales(value['Reporter'], value['Count']));
    }

    return [
      charts.Series<Sales, String>(
        id: 'Sales',
        domainFn: (Sales sales, _) => sales.x,
        measureFn: (Sales sales, _) => sales.y,
        data: desktopSalesData,
        fillColorFn: (Sales sales, _) {
          return charts.MaterialPalette.blue.shadeDefault;
        },
      )
    ];
  }

  Future<void> _fetchData() async {
    var seriesList = await _createRandomData();

    setState(() {
      this.seriesList = seriesList;
    });
  }

  @override
  void initState() {
    super.initState();

    _fetchData();
  }

  barChart() {
    return charts.BarChart(
      seriesList,
      animate: true,
      vertical: true,
      barGroupingType: charts.BarGroupingType.grouped,
      defaultRenderer: charts.BarRendererConfig(
        groupingType: charts.BarGroupingType.grouped,
        strokeWidthPx: 1.0,
      ),
      domainAxis: const charts.OrdinalAxisSpec(
      renderSpec: charts.SmallTickRendererSpec(labelRotation: 60,
      labelStyle: charts.TextStyleSpec(
          fontSize: 12,)
    ),)
    );
  }

  Future<List<dynamic>> kpi3() async {
    var kpi = Uri.parse("http://172.18.0.1:8080/employees/kpi3");

    var res = await http.get(kpi);

    if (res.statusCode == 200) {
      // Request successful

      var jsonData = jsonDecode(res.body);

      return jsonData;
    } else {
      // Request failed

      throw Exception('Unxexpected error occured');
    }
  }

  Future<List<dynamic>> kpi4() async {
    var kpi = Uri.parse("http://172.18.0.1:8080/employees/kpi4");

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
    var kpi = Uri.parse("http://172.18.0.1:8080/employees/kpi6");

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
                width: screenWidth * 0.9,
                height: screenHeight * 0.4,
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
              ),
              const SizedBox(
                height: 20,
              ),
              Container(
                width: screenWidth * 0.9,
                height: screenHeight * 0.4,
                decoration: BoxDecoration(
                  borderRadius: BorderRadius.circular(12),
                  color: const Color(0xFFFFFFFF).withOpacity(0.50),
                ),
                child: barChart(),
              ),
              const SizedBox(
                height: 20,
              ),
              FutureBuilder<List<dynamic>>(
              future: kpi6(),
              builder: (context, snapshot) {
                if (snapshot.connectionState == ConnectionState.waiting) {
                  return const Center(child: CircularProgressIndicator());
                } else if (snapshot.hasError) {
                  // ... Error handling ...
                  return Center(child: Text('Error: ${snapshot.error}'));
                } else {
                  allTeamsDataMap = dataMap(snapshot.data ?? []);
                  return Column(
                    children: allTeamsDataMap.entries.map((entry) {
                      String teamName = entry.key;
                      Map<String, double> teamDataMap = entry.value;

                      return Column(
                        children: [
                          Container(
                            width: screenWidth * 0.9,
                            height: screenHeight * 0.4,
                            decoration: BoxDecoration(
                              borderRadius: BorderRadius.circular(12),
                              color: const Color(0xFFFFFFFF).withOpacity(0.50),
                            ),
                            child: piechart.PieChart(
                              dataMap: teamDataMap,
                              // ... Other pie chart options ...
                            ),
                          ),
                          const SizedBox(height: 10),
                          Text(
                            teamName,
                            style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold),
                          ),
                        ],
                      );
                    }).toList(),
                  );
                }
              },
            ),

            ],
          ),
        ),
      ),
    );
  }
}



