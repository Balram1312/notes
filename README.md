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
  const KpiDashboard({Key? key}) : super(key: key);

  @override
  _KpiDashboardState createState() => _KpiDashboardState();
}

class _KpiDashboardState extends State<KpiDashboard> {
  List<charts.Series<Sales, String>> seriesList = [];
  Map<String, Map<String, double>> allTeamsDataMap = {};

  @override
  void initState() {
    super.initState();
    _fetchData();
  }

  Future<void> _fetchData() async {
    var seriesList = await _createRandomData();
    setState(() {
      this.seriesList = seriesList;
    });
  }

  Future<List<charts.Series<Sales, String>>> _createRandomData() async {
    var values = await kpi4();
    List<Sales> desktopSalesData = values.map((value) {
      return Sales(value['Reporter'], value['Count']);
    }).toList();
    return [
      charts.Series<Sales, String>(
        id: 'Sales',
        domainFn: (Sales sales, _) => sales.x,
        measureFn: (Sales sales, _) => sales.y,
        data: desktopSalesData,
        fillColorFn: (Sales sales, _) =>
            charts.MaterialPalette.blue.shadeDefault,
      ),
    ];
  }

  Future<List<dynamic>> kpi3() async {
    var kpi = Uri.parse("http://172.18.0.1:8080/employees/kpi3");
    var res = await http.get(kpi);

    if (res.statusCode == 200) {
      var jsonData = jsonDecode(res.body);
      return jsonData;
    } else {
      throw Exception('Unexpected error occurred');
    }
  }

  Future<List<dynamic>> kpi4() async {
    var kpi = Uri.parse("http://172.18.0.1:8080/employees/kpi4");
    var res = await http.get(kpi);

    if (res.statusCode == 200) {
      var jsonData = jsonDecode(res.body);
      return jsonData;
    } else {
      throw Exception('Failed to fetch data');
    }
  }

  Future<List<dynamic>> kpi6() async {
    var kpi = Uri.parse("http://172.18.0.1:8080/employees/kpi6");
    var res = await http.get(kpi);

    if (res.statusCode == 200) {
      var jsonData = jsonDecode(res.body);
      return jsonData;
    } else {
      throw Exception('Failed to fetch data');
    }
  }

  Map<String, Map<String, double>> dataMap(List<dynamic> kpi6Data) {
    Map<String, Map<String, double>> teamDataMap = {};

    for (var entry in kpi6Data) {
      String teamName = entry['Teamname'];
      String leaveType = entry['Leavetype'];
      double leaveCount = entry['Leavecount'].toDouble();

      teamDataMap.putIfAbsent(teamName, () => {});
      teamDataMap[teamName]![leaveType] = leaveCount;
    }

    return teamDataMap;
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
        renderSpec: charts.SmallTickRendererSpec(
          labelRotation: 60,
          labelStyle: charts.TextStyleSpec(
            fontSize: 12,
          ),
        ),
      ),
    );
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
          icon: const Icon(Icons.arrow_left),
        ),
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
                    if (snapshot.connectionState == ConnectionState.waiting) {
                      return const Center(child: CircularProgressIndicator());
                    } else if (snapshot.hasError) {
                      return Text('Error: ${snapshot.error}');
                    } else {
                      List<dynamic> data = snapshot.data ?? [];
                      List<DataRow> rows = data.map((data) {
                        return DataRow(cells: [
                          DataCell(Text(data['Name'] ?? '')),
                          DataCell(Text(data['Totalleave2023'].toString())),
                        ]);
                      }).toList();

                      return SingleChildScrollView(
                        child: DataTable(
                          columns: const [
                            DataColumn(label: Text('NAME')),
                            DataColumn(label: Text('MOST LEAVE 2023')),
                          ],
                          rows: rows,
                        ),
                      );
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
                                color:
                                    const Color(0xFFFFFFFF).withOpacity(0.50),
                              ),
                              child: piechart.PieChart(
                                dataMap: teamDataMap,
                                // ... Other pie chart options ...
                              ),
                            ),
                            const SizedBox(height: 10),
                            Text(
                              teamName,
                              style: const TextStyle(
                                  fontSize: 18, fontWeight: FontWeight.bold),
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



---------------------------------------
reg page
---------------------------------------
import 'package:flutter/material.dart';
import 'package:wc_form_validators/wc_form_validators.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'package:fluttertoast/fluttertoast.dart';

class RegistrationPage extends StatefulWidget {
  const RegistrationPage({super.key});

  @override
  State<RegistrationPage> createState() => _RegistrationPageState();
}

class _RegistrationPageState extends State<RegistrationPage> {
  bool isChecked = false;
  final formKey = GlobalKey<FormState>();
  late String labelText;

  final TextEditingController leavetypeController = TextEditingController();
  final TextEditingController teamnameController = TextEditingController();
  final TextEditingController reporterController = TextEditingController();
  final TextEditingController fromdateController = TextEditingController();
  final TextEditingController todateController = TextEditingController();
final TextEditingController nameController = TextEditingController();

  void registerUser() async {
    var url = Uri.parse('http://172.18.0.1:8080/employee');
    var headers = {'Content-Type': 'application/json'};

    var jsonData = {
      "name": nameController.text,
      "leavetype": leavetypeController.text,
      "fromdate": fromdateController.text,
      "todate": todateController.text,
      "teamname": teamnameController.text,
      "reporter": reporterController.text,
    };

    debugPrint('hello $jsonData');
    var body = jsonEncode(jsonData);
    var response = await http.post(url, headers: headers, body: body);

    if (response.statusCode == 200) {
      Fluttertoast.showToast(
        msg: 'REGISTERED SUCCESSFULLY !!',
        toastLength: Toast.LENGTH_SHORT,
        gravity: ToastGravity.CENTER,
      );
      debugPrint(response.body);
    } else {
      // Request failed
      debugPrint('Request failed with status: ${response.statusCode}');
      debugPrint(response.body);
    }
  }

  Future<void> _selectFromDate(BuildContext context) async {
    final DateTime? picked = await showDatePicker(
      context: context,
      initialDate: DateTime.now(),
      firstDate: DateTime(2000),
      lastDate: DateTime(2101),
    );
    if (picked != null && picked != DateTime.now()) {
      setState(() {
        fromdateController.text = picked.toLocal().toString().split(' ')[0];
      });
    }
  }

  Future<void> _selectToDate(BuildContext context) async {
    final DateTime? picked = await showDatePicker(
      context: context,
      initialDate: DateTime.now(),
      firstDate: DateTime(2000),
      lastDate: DateTime(2101),
    );
    if (picked != null && picked != DateTime.now()) {
      setState(() {
        todateController.text = picked.toLocal().toString().split(' ')[0];
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    double screenWidth = MediaQuery.of(context).size.width;
    double screenHeight = MediaQuery.of(context).size.height;
    return Scaffold(
      backgroundColor: const Color(0xFF009688),
      appBar: AppBar(
        centerTitle: true,
        backgroundColor: Colors.teal,
        elevation: 0,
        leading: IconButton(
          onPressed: () {
            Navigator.pop(context);
          },
          icon: const Icon(Icons.arrow_left),
        ),
        title: const Text(
          'REGISTRATION',
          style: TextStyle(fontSize: 20, letterSpacing: 1.5),
        ),
      ),
      body: SingleChildScrollView(
        child: Form(
          key: formKey,
          child: Center(
            child: Container(
              padding: const EdgeInsets.all(20),
              width: screenWidth * 0.9,
              height: screenHeight * 0.85,
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(12),
                color: const Color(0xFFFFFFFF).withOpacity(0.25),
              ),
              child: Column(
                mainAxisAlignment: MainAxisAlignment.spaceAround,
                children: [
                  const Image(
                    image: AssetImage('images/bg_second_page.png'),
                    color: Colors.white,
                    width: 100,
                  ),
                  TextFormField(
                    controller: nameController,
                    style: buildTextStyleFormFields(),
                    validator: Validators.required('Name is required'),
                    decoration: buildInputDecoration('Name'),
                  ),
                  DropdownButtonFormField(
                    
                   
                    onChanged: (value) {
                      leavetypeController.text = value.toString();
                      
                    },
                    items: const [
                      DropdownMenuItem(
                        value: 'Casual Leave',
                        child: Text('Casual Leave'),
                      ),
                      DropdownMenuItem(
                        value: 'Earned Leave',
                        child: Text('Earned Leave'),
                      ),
                      DropdownMenuItem(
                        value: 'Sick Leave',
                        child: Text('Sick Leave'),
                      ),
                    ],
                    validator: Validators.required('Reporter is required'),
                    decoration: buildInputDecoration('Leave Type'),
                  ),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      
                      Expanded(
                        child: GestureDetector(
                                          onTap: () => _selectFromDate(context),
                                          child: AbsorbPointer(
                        child: TextFormField(
                          controller: fromdateController,
                          style: buildTextStyleFormFields(),
                          validator: Validators.required('From Date is required'),
                          decoration: buildInputDecoration('From Date'),
                        ),
                                          ),
                                        ),
                      ),
                  Expanded(
                    child: GestureDetector(
                      onTap: () => _selectToDate(context),
                      child: AbsorbPointer(
                        child: TextFormField(
                          controller: todateController,
                          style: buildTextStyleFormFields(),
                          validator: Validators.required('To Date is required'),
                          decoration: buildInputDecoration('To Date'),
                        ),
                      ),
                    ),
                  ),
                    ],
                  ),
                  
                  Row(
                    children: [
                      const Text(
                        'Team Name:',
                        style: TextStyle(
                          fontSize: 18,
                          color: Colors.white,
                        ),
                      ),
                      const SizedBox(width: 10),
                      Expanded(
                        child: Column(
                          children: [
                            RadioListTile<String>(
                              activeColor: Colors.white,
                              fillColor: MaterialStateColor.resolveWith((states) => Colors.white),
                              title: const Text(
                                'AWS',
                                style: TextStyle(
                                  color: Colors.white,
                                ),
                              ),
                              value: 'aws',
                              groupValue: teamnameController.text,
                              onChanged: (String? value) {
                                setState(() {
                                  teamnameController.text = value!;
                                });
                              },
                            ),
                            RadioListTile<String>(
                              activeColor: Colors.white,
                              fillColor: MaterialStateColor.resolveWith((states) => Colors.white),

                              title: const Text(
                                'Azure',
                                style: TextStyle(
                                  color: Colors.white,
                                ),
                              ),
                              value: 'azure',
                              groupValue: teamnameController.text,
                              onChanged: (String? value) {
                                setState(() {
                                  teamnameController.text = value!;
                                });
                              },
                            ),
                          ],
                        ),
                      ),
                    ],
                  ),
                  DropdownButtonFormField(
                    
                   
                    onChanged: (value) {
                      reporterController.text = value.toString();
                      
                    },
                    items: const [
                      DropdownMenuItem(
                        value: 'john hopkins',
                        child: Text('John Hopkins'),
                      ),
                      DropdownMenuItem(
                        value: 'raj singh',
                        child: Text('Raj Singh'),
                      ),
                      DropdownMenuItem(
                        value: 'pooja mishra',
                        child: Text('Pooja Mishra'),
                      ),
                    ],
                    validator: Validators.required('Reporter is required'),
                    decoration: buildInputDecoration('Reporter'),
                  ),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.end,
                    children: [
                      Checkbox(
                        checkColor: const Color(0xFF009866),
                        fillColor: const MaterialStatePropertyAll(Colors.white),
                        value: isChecked,
                        onChanged: (bool? value) {
                          setState(() {
                            isChecked = value!;
                          });
                        },
                      ),
                      const Text(
                        'Agree with Terms & Condition.',
                        style: TextStyle(
                          color: Color(0xFFFFFFFF),
                          fontWeight: FontWeight.bold,
                        ),
                      )
                    ],
                  ),
                  ElevatedButton(
                    onPressed: () {
                      if (isChecked && formKey.currentState!.validate()) {
                        // Checkbox is checked and form data is valid
                        registerUser();
                        formKey.currentState?.reset();
                        leavetypeController.clear();
                        fromdateController.clear();
                        todateController.clear();
                        teamnameController.clear();
                        reporterController.clear();
                      } else {
                        Fluttertoast.showToast(
                          msg: 'Please agree to the Terms & Conditions',
                          toastLength: Toast.LENGTH_SHORT,
                          gravity: ToastGravity.CENTER,
                        );
                      }
                    },
                    style: ElevatedButton.styleFrom(
                      elevation: 0,
                      backgroundColor: const Color(0xFF009688),
                      minimumSize: Size(screenWidth * .9, screenHeight * .06),
                    ),
                    child: const Text('Submit'),
                  ),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }

  TextStyle buildTextStyleFormFields() {
    return const TextStyle(color: Colors.white);
  }

  InputDecoration buildInputDecoration(String labelText) {
    return InputDecoration(
      fillColor: Colors.teal[50],
      labelText: labelText,
      labelStyle: const TextStyle(fontSize: 18, color: Colors.white),
      focusedBorder: const OutlineInputBorder(
        borderSide: BorderSide(color: Colors.white, width: 1.8),
      ),
      enabledBorder: const OutlineInputBorder(
        borderSide: BorderSide(color: Colors.white, width: 3),
      ),
      errorBorder: const OutlineInputBorder(
        borderSide: BorderSide(color: Colors.red, width: 1),
      ),
    );
  }
}

----------------------------------------------------
k page
----------------------------------------------------

import 'dart:convert';

import 'dart:async';

import 'package:charts_flutter/flutter.dart' as charts;

import 'package:flutter/material.dart';

import 'package:http/http.dart' as http;

import 'package:pie_chart/pie_chart.dart' as piechart;

class KpiDashboard extends StatefulWidget {
  const KpiDashboard({Key? key}) : super(key: key);

  @override
  _KpiDashboardState createState() => _KpiDashboardState();
}

class _KpiDashboardState extends State<KpiDashboard> {
  List<charts.Series<Sales, String>> seriesList = [];

  Map<String, Map<String, double>> allTeamsDataMap = {};

  String endpointUrl = "http://172.18.0.1:8080/employees/";

  @override
  void initState() {
    super.initState();

    _fetchData();
  }

  Future<void> _fetchData() async {
    var seriesList = await _createRandomData();

    setState(() {
      this.seriesList = seriesList;
    });
  }

  Future<List<charts.Series<Sales, String>>> _createRandomData() async {
    var values = await kpi4();

    List<Sales> desktopSalesData = values.map((value) {
      return Sales(value['Reporter'], value['Count']);
    }).toList();

    return [
      charts.Series<Sales, String>(
        id: 'Sales',
        domainFn: (Sales sales, _) => sales.x,
        measureFn: (Sales sales, _) => sales.y,
        data: desktopSalesData,
        fillColorFn: (Sales sales, _) =>
            charts.MaterialPalette.blue.shadeDefault,
      ),
    ];
  }

  Future<List<dynamic>> kpi3() async {
    var kpi = Uri.parse("${endpointUrl}kpi3");

    var res = await http.get(kpi);

    if (res.statusCode == 200) {
      var jsonData = jsonDecode(res.body);

      return jsonData;
    } else {
      throw Exception('Unexpected error occurred');
    }
  }

  Future<List<dynamic>> kpi4() async {
    var kpi = Uri.parse("${endpointUrl}kpi4");

    var res = await http.get(kpi);

    if (res.statusCode == 200) {
      var jsonData = jsonDecode(res.body);

      return jsonData;
    } else {
      throw Exception('Failed to fetch data');
    }
  }

  Future<List<dynamic>> kpi6() async {
    var kpi = Uri.parse("${endpointUrl}kpi6");

    var res = await http.get(kpi);

    if (res.statusCode == 200) {
      var jsonData = jsonDecode(res.body);

      return jsonData;
    } else {
      throw Exception('Failed to fetch data');
    }
  }

  Map<String, Map<String, double>> dataMap(List<dynamic> kpi6Data) {
    Map<String, Map<String, double>> teamDataMap = {};

    for (var entry in kpi6Data) {
      String teamName = entry['Teamname'];

      String leaveType = entry['Leavetype'];

      double leaveCount = entry['Leavecount'].toDouble();

      teamDataMap.putIfAbsent(teamName, () => {});

      teamDataMap[teamName]![leaveType] = leaveCount;
    }

    return teamDataMap;
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
        renderSpec: charts.SmallTickRendererSpec(
          labelRotation: 60,
          labelStyle: charts.TextStyleSpec(
            fontSize: 12,
          ),
        ),
      ),
    );
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
          icon: const Icon(Icons.arrow_back),
        ),
        title: const Text(
          'KPI DASHBOARD',
          style: TextStyle(fontSize: 20, letterSpacing: 1.5),
        ),
        actions: [
          IconButton(
            onPressed: _refreshAction,
            icon: const Icon(Icons.refresh),
            color: Colors.white,
          )
        ],
      ),
      body: SingleChildScrollView(
        child: Center(
          child: Column(
            children: [
              TitleContainer(
                screenWidth: screenWidth,
                title: 'Employees With Maximum Leave Taken.',
              ),
              Container(
                width: screenWidth * 0.9,
                decoration: BoxDecoration(
                  borderRadius: BorderRadius.circular(12),
                  color: const Color(0xFFFFFFFF).withOpacity(0.50),
                ),
                child: FutureBuilder<List<dynamic>>(
                  future: kpi3(),
                  builder: (context, snapshot) {
                    if (snapshot.connectionState == ConnectionState.waiting) {
                      return const Center(child: CircularProgressIndicator());
                    } else if (snapshot.hasError) {
                      return Text('Error: ${snapshot.error}');
                    } else {
                      List<dynamic> data = snapshot.data ?? [];

                      List<DataRow> rows = data.map((data) {
                        return DataRow(cells: [
                          DataCell(Text(data['Name'] ?? '')),
                          DataCell(Text(data['Totalleave2023'].toString())),
                        ]);
                      }).toList();

                      return SingleChildScrollView(
                        child: DataTable(
                          columns: const [
                            DataColumn(label: Text('NAME')),
                            DataColumn(label: Text('MOST LEAVE 2023')),
                          ],
                          rows: rows,
                        ),
                      );
                    }
                  },
                ),
              ),
              const SizedBox(
                height: 20,
              ),
              TitleContainer(
                screenWidth: screenWidth,
                title: 'Employees Leaves Distribution Under Each Manager.',
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
              TitleContainer(
                screenWidth: screenWidth,
                title: 'Distribution of Most Leave Taken by 2 Teams',
              ),
              FutureBuilder<List<dynamic>>(
                future: kpi6(),
                builder: (context, snapshot) {
                  if (snapshot.connectionState == ConnectionState.waiting) {
                    return const Center(child: CircularProgressIndicator());
                  } else if (snapshot.hasError) {
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
                              margin: const EdgeInsets.only(bottom: 5),
                              width: screenWidth * 0.9,
                              decoration: BoxDecoration(
                                borderRadius: BorderRadius.circular(12),
                                color:
                                    const Color(0xFFFFFFFF).withOpacity(0.50),
                              ),
                              child: piechart.PieChart(
                                dataMap: teamDataMap,
                                centerText: teamName,
                              ),
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

  _refreshAction() {
    setState(() {});
  }
}

class TitleContainer extends StatelessWidget {
  const TitleContainer({
    super.key,
    required this.screenWidth,
    required this.title,
  });

  final String title;

  final double screenWidth;

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.all(5),
      margin: const EdgeInsets.only(bottom: 5),
      width: screenWidth * 0.9,
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(12),
        color: const Color(0xFFFFFFFF).withOpacity(0.50),
      ),
      child: Text(
        title,
        textAlign: TextAlign.center,
        style: const TextStyle(fontWeight: FontWeight.bold, color: Colors.teal),
      ),
    );
  }
}

class Sales {
  final String x;

  final int y;

  Sales(this.x, this.y);
}

-------------------------------------------------------
/*


Conclusion: Unleashing the Power of Amazon Pinpoint
In conclusion, Amazon Pinpoint emerges as a transformative force in the realm of customer communication and engagement. Throughout this presentation, we've explored the core features and functionalities that make Amazon Pinpoint a crucial asset for businesses seeking to elevate their outreach strategies.

Amazon Pinpoint's multi-channel messaging capabilities, coupled with personalization and targeting, offer a dynamic approach to reaching customers across diverse platforms. The robust analytics and insights empower businesses with data-driven decision-making, ensuring campaigns are not just sent but are strategically optimized for success.

Automation features, user segmentation, and A/B testing contribute to the agility and effectiveness of communication strategies. Whether it's automating triggered campaigns, tailoring messages for specific audience segments, or experimenting with different approaches through A/B testing, Pinpoint empowers businesses to refine and tailor their messaging with precision.

Furthermore, the seamless integration with other AWS services expands the horizons of possibility, allowing businesses to harness the collective strength of Amazon's cloud infrastructure.

As we navigate the dynamic landscape of customer engagement, Amazon Pinpoint stands as a reliable ally, offering the tools necessary to craft compelling narratives, foster meaningful connections, and drive results. As we embark on the journey of leveraging Amazon Pinpoint, let's embrace the opportunity to revolutionize how we communicate with our audience, creating experiences that resonate and endure. Together, we unlock the potential to not only meet but exceed the expectations of our customers, fostering lasting relationships and driving success in our communication endeavors.
*/







balram+1-at-390431833332,+bWSnZMmAvXgxXiNnevZohbxQBBwJJMW5fGuW21I4zg=






