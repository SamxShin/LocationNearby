# LocationNearby

This project implements google's Geocode API and allows the user to either enter an address containing the street number, street name, city name, state name, and country name (ie 123 Fake Street, Los Angeles CA, USA) and would return the latitude and longitude.
There is also reverse Geocode option where the user can give longitude and latitude, and the program will return the address according to those components.
These will be in the LocationNearvy Folder

Limitations: The original plan was to use a relational database such as PostgreSQL and store the locations using tables with each of the different address components, but to limited experience using databases, I was not able to successfully implement those components into the main program. However, I was able to make a separate program that can store those data values in to PostgreSQL, however, it is a very naive program. This program has a for-loop that allows the user to put in 3 places manually and stores those values into PostgreSQL database. For a future implementation, I will have this in the main program to go along with the point-break loop.
The naive program and picture of output into PostgreSQL database will be in the database folder
*please ignore the data inputted in the image folder, that picutre was to show for testing purposes only*

Thanks, Samuel Shin
