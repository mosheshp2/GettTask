# GettTask
Gett GO lang task solution

This is a solution for the Gett task,

1. I created 2 tables in Postgres SQL, "Metrics", and "Drivers"

2. To upload the json files data I created the PKG "fileread", which exposes to functions, "GetDrivers", and "ProcessMetrics"
   Drivers are loaded to array of struct "Driver" 
   Metrics are loaded ad text lines.

   To import data to SQL I created the sqlimport.go that compiled to sqlimport.exe, that uses the fileread pkg
   we take the arrays from fileread, and then iterate of the arrays and inserts them to SQL.
   
3. For the third task, I used "Beego" framework, to create skeleton project of REST api for generic collections
   I ddded support to the Drivers collection, BUT DIDN'T finished the whole CRUD operation due to lack of time...
   
Hope it's good enough

Moshe Shpitz
