                                    LEARN GO, GO LEARN


Week 1:


Why?

•	It is a strong and statically typed language means each variable has a type.

•	Go is designed for faster compilation without the need for dependency checking. No libraries or framework is  used, so no dependency.

•	It has a lot of string functions to manipulate string, also it has concept of garbage collection.

•	Concurrency: It is related to multitasking, means whenever any memory gets free, the next task will be executed. The task is not dependent to each other.

           GO = C + strings + garbage collection + concurrency

•	Products or tools on Go: Docker, Kubernetes, Ethereum etc.


* If we have to use different packages in different files, then we have to keep them in different folder. It is because Go follows strict file structure. It means we can use only one package in a folder of different files. 
Refer my_package folder for above.

* If the first letter of a variable name is capital , then we can access it in other packages. Same is for functions.

* User Input: scan() reads till first whitespace character. Generally we use 'bufio' package for user input. A buffered reader is a type of reader that reads data from underlying source, such as a file or keyboard, and stores that data in a buffer. The buffer is a temporary storage area in memory. It is commonly used to improve the efficiency of input operations.

* _ is a blank and used to ignore the return data of a function, because in go whataver we are declaring, we have to use it in our program. In some cases we don't want to use them, then we store them in _.

* slice is a flexible and dynamic data structure. slice <---> vector.

* map is just like hashing which stores key value pairs and it is unordered.

* struct: It is a composite data type that groups together variables under a single name.

* pointers: they provide a way to work with memory directly which helps in efficient memory management and sharing data between functions

* time: 02-01-2006, Monday. Time function in go used this specific format. For formatting time, we use format function.

* File Handling: When you open a file using functions like os.Create(), os.Open(), you are acquiring system resources to interact with that file. Failing to close the file properly through file.Close() can lead to resource leaks and might cause issues like running out of file descriptors.

* Web Request: Don't forget to close the body with defer keyword after getting the data from any api. Refer web_request folder.

* URL handling: Convert the string url through url.Parse() method, then take the data.

* CRUD: To get the data from api , we prefer to make structure for the data we want to get through api. Refer CRUD folder.

* Goroutine: It allow you to run functions concurrently, or in parallel, with other parts of your program. It improves the performance of the program, especially when dealing with tasks that involve I/O operations or other types of blocking operations.

* Project: URL SHORTNER

1- We have made a structure/model, means these are the data which are stored in database.

2- Made a database using map data structure, means the short url will be mapped with the data which are defined in model

NOTE: Whenever some data comes from the web, it comes in the form of byte, we have to convert it into string

*****************CHECK PRACTICE FOLDER, YOU WILL SEE SOMETHING INTERESTING*******************************