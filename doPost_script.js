function doPost(e) {
  const sheet = SpreadsheetApp.getActiveSpreadsheet().getSheetByName("Sheet1");
  const data = JSON.parse(e.postData.contents);
  
  sheet.appendRow([
    new Date(),
    data.lat,
    data.lng,
    data.azimuth,
    data.fov,
    data.height,
    data.tilt
  ]);

  return ContentService.createTextOutput("Success");
}

