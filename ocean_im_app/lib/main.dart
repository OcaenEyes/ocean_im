import 'dart:io';

import 'package:flutter/material.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:ocean_im/layout/auth/login.dart';
import 'package:ocean_im/router/page_builder.dart';
import 'package:ocean_im/router/page_router.dart';
import 'package:system_tray/system_tray.dart';
import 'package:window_manager/window_manager.dart';

void main() async {
  PageRouter.setupRoutes();
  WidgetsFlutterBinding.ensureInitialized();
  await windowManager.ensureInitialized();

  windowManager.waitUntilReadyToShow().then(((value) async => {
        // await windowManager.setTitleBarStyle('hidden'),
        await windowManager.setMinimumSize(const Size(800, 600)),
        await windowManager.setSize(const Size(800, 600)),
        await windowManager.setResizable(true),
        // await windowManager.setMovable(true),
        await windowManager.setMinimizable(true),
        await windowManager.setClosable(true),
        await windowManager.center(),
        await windowManager.show(),
        await windowManager.setSkipTaskbar(false)
      }));

  // runApp(MultiProvider(providers: [
  //   ChangeNotifierProvider(create: (BuildContext context) => MessageProvider()),
  //   ChangeNotifierProvider(
  //       create: (BuildContext context) => UserInfoNotifier()),
  // ], child: const MyApp()));
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
        title: 'OCEAN IM',
        builder: EasyLoading.init(),
        onGenerateRoute: PageRouter.router.generator,
        theme: ThemeData(
            // This is the theme of your application.
            //
            // Try running your application with "flutter run". You'll see the
            // application has a blue toolbar. Then, without quitting the app, try
            // changing the primarySwatch below to Colors.green and then invoke
            // "hot reload" (press "r" in the console where you ran "flutter run",
            // or simply save your changes to "hot reload" in a Flutter IDE).
            // Notice that the counter didn't reset back to zero; the application
            // is not restarted.
            // primarySwatch: Colors.blue,
            primaryColor: Colors.black,
            inputDecorationTheme: const InputDecorationTheme(
                iconColor: Colors.black,
                focusedBorder: UnderlineInputBorder(
                    borderSide: BorderSide(color: Colors.black)))),
        home: const MyHomePage());
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({Key? key}) : super(key: key);

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  var isLogin = false;
  final SystemTray systemTray = SystemTray();
  String iconpath = Platform.isWindows
      ? 'assets/ico/app_icon.ico'
      : 'assets/ico/app_icon.png';

  @override
  void initState() {
    // getLocalData();
    super.initState();
    initSystemTray();
  }

  // dynamic getLocalData() async {
  //   SharedPreferences sharedPreferences = await SharedPreferences.getInstance();
  //   var localData = sharedPreferences.getString("userInfo").toString();
  //   if (localData != "null") {
  //     setState(() {
  //       isLogin = true;
  //     });
  //   }
  // }

  Future<void> initSystemTray() async {
    final menu = [
      MenuItem(
          label: "显示",
          onClicked: () {
            windowManager.show();
          }),
      MenuItem(
          label: "隐藏",
          onClicked: () {
            windowManager.hide();
          }),
      MenuItem(
          label: "关闭",
          onClicked: () {
            windowManager.close();
          }),
    ];
    await systemTray.initSystemTray(
        title: Platform.isWindows ? "菜单" : "",
        iconPath: iconpath,
        toolTip: "OCEAN-IM");

    await systemTray.setContextMenu(menu);
    systemTray.registerSystemTrayEventHandler((eventName) {
      if (eventName == "rightMouseUp") {
        systemTray.popUpContextMenu();
      } else if (eventName == "leftMouseUp") {
        windowManager.show();
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    return Login(bundleArguments: BundleArguments());
  }
}
