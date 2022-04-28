class DatabaseUtil {
  static final DatabaseUtil databaseUtil = DatabaseUtil._instance();
  DatabaseUtil._instance();

  // late Database database;
  // Future<Database> get db async {
  //   if (database != null) {
  //     return database;
  //   }
  //   database = await getDatabaseInstance();
  //   return database;
  // }

  // Future<Database> getDatabaseInstance() async {
  //   var dbFilePath = '${FileUtil.fileUtil.localDocumentPath}' "/databases";
  //   FileUtil.fileUtil.createFileFolder(dbFilePath);
  //   String dbPath = join(dbFilePath, "im.db");

  //   return sqlite3.open(dbPath);
  // }
}
