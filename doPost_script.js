function doPost(e) {
  // Add sheet name here
  const sheet = SpreadsheetApp.getActiveSpreadsheet().getSheetByName("camera_map_data");
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

