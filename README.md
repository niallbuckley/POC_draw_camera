# Proof of concept: accurately define geo-location and view coordinates for a camera and export to csv

## Set up - Simple
Run command \
`make app`

## Set up - Complex
Run the HTML in a page which does not have web secuirty (CORS will typically block this request)\
`google-chrome --disable-web-security --user-data-dir=~/Desktop/test/`

You can write to the sheet referenced in the code if you want, or else edit your own sheet

To edit your own sheet:
- Open your Google Sheet
- Click Extensions > Apps Script
- Add the post script
- Save the project
- Click Deploy > Manage deployments
- Click "New deployment"
- Choose "Web App"
Execute as: Me
Access: Anyone (or "Anyone with the link")

- Click Deploy
- Copy the Web App URL and add to camera_map.html
