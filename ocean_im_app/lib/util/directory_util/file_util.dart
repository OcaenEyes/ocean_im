import 'dart:io';

import 'package:flutter/cupertino.dart';
import 'package:path_provider/path_provider.dart';

class FileUtil {
  static FileUtil fileUtil = FileUtil._instance();

  FileUtil._instance();

  get localDocumentPath async {
    var directory = await getApplicationDocumentsDirectory();
    return directory.path;
  }

  get localTempPath async {
    var directory = await getTemporaryDirectory();
    return directory.path;
  }

  get localSupportPath async {
    var directory = await getApplicationSupportDirectory();
    return directory.path;
  }

//创建文件夹
  createFileFolder(String path) async {
    Directory directory = Directory(path);
    if (!directory.existsSync()) {
      directory.createSync(recursive: true);
      debugPrint("文件夹新建成功， ${directory.path}");
    } else {
      debugPrint("文件夹已存在");
    }

    return directory.path;
  }

// 创建文件
  createFile(String path) async {
    File file = File(path);
    if (!file.existsSync()) {
      file.createSync(recursive: true);
      debugPrint("文件新建成功， ${file.path}");
    } else {
      debugPrint("文件已存在");
    }

    return file;
  }

//写入文件
  writeFile(String path, String content) async {
    File file = await createFile(path);
    try {
      file.writeAsStringSync(content);
      debugPrint('文件写入成功');
    } catch (e) {
      debugPrint('文件写入失败');
    }
  }

  // 读取文件
  readFile(String path) async {
    File file = await createFile(path);
    try {
      String contents = file.readAsStringSync();
      debugPrint('文件读取成功');
      return contents;
    } catch (e) {
      debugPrint('文件读取失败');
      return "";
    }
  }

  ///文件/文件夹删除
  deleteFilefloder(String path) {
    Directory directory = Directory(path);
    if (directory.existsSync()) {
      List<FileSystemEntity> files = directory.listSync();
      if (files.isNotEmpty) {
        for (var file in files) {
          file.deleteSync();
        }
      }
      directory.deleteSync();
      debugPrint('文件夹删除成功');
    }
  }

  ///删除文件
  deleteFile(String path) {
    File file = File(path);
    if (file.existsSync()) {
      try {
        file.deleteSync();
        debugPrint('文件删除成功');
      } catch (e) {
        debugPrint('文件删除失败');
      }
    }
  }

  //遍历所有文件
  getAllSubFile(String path) async {
    Directory directory = Directory(path);
    if (directory.existsSync()) {
      return directory.listSync();
    }
    return [];
  }
}
