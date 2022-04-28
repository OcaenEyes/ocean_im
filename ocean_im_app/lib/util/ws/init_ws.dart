import 'package:ocean_im/common/global_info.dart';
import 'package:ocean_im/util/env.dart';
import 'package:socket_io_client/socket_io_client.dart' as IO;

IO.Socket initWs() {
  IO.Socket socket = IO.io(
      devBaseUrl,
      IO.OptionBuilder().setPath("/ws").disableAutoConnect().setQuery(
          {"token": Global.token}).setTransports(["websocket"]).build());
  return socket;
}
