import 'package:ocean_im/layout/auth/login.dart';
import 'package:ocean_im/layout/view/main_view/main_view.dart';
import 'package:ocean_im/router/page_builder.dart';

enum PageName { login, mainView }

Map<PageName, PageBuilder> pageRoutes = {
  PageName.login:
      PageBuilder(pageBuilderFunc: (bundle) => Login(bundleArguments: bundle)),
  PageName.mainView: PageBuilder(
      pageBuilderFunc: (bundle) => MainView(bundleArguments: bundle)),
};
