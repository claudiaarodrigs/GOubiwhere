Projecto GO - UBIWHERE

The goal of this project was to accomplish a Go application that was able to collect CPU and RAM usage each second and store it in a local database, create a simulator of an external device generating (each second) a sample composed by four variables and store them in a local database, and, finally, to provide an interface between the user and the samples data through the console allowing multiple processing from that data.
To allow the coexistence between the database continuous fill, and, at the same time, its interaction, it was created two separated apps: DataCollector.exe and Interface.exe, available in /bin directory of this repository. 
To execute, both apps must be in the same directory, and run the DataCollector in one terminal, and, without closing this run, use the Interface app in other terminal.
