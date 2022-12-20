# SubscribeMe

Application for tracking tasks within one click for students coded in Flutter and Golang.

<img width="214" height="400" alt="iPhone 11 Pro Max - 1" src="https://user-images.githubusercontent.com/54924480/204043212-ff4c3a2f-8c81-4903-96c3-46d11720c78f.png">


## Background
During my college career, I often saw many students forget their homework and submit it over the deadline. This application helps them track their homework, so they can prepare to finish it within time.

## Features
### 1. Register
<img width="214" height="400" alt="register" src="https://user-images.githubusercontent.com/54924480/208782092-ded6ff19-3793-43b6-bd38-5c8e58a37935.png">
Users can register by entering their username and password. This feature use JWT-based authenticatoin using access token and refresh token. Both of the tokens are stored in local device using Flutter Secure Storage. The password is hashed using bycript.

### 2. Login
<img width="214" height="400" alt="login" src="https://user-images.githubusercontent.com/54924480/208781878-cf7d77f4-2ed1-4bdb-ae42-04bfd4f7598a.png">
The application will check whether the access token or refresh token is expired. If not, the application will automatically log in the user. If yes, the application will prompt the user to enter their username and password.

### 3. View today deadlines and deadline 7 days ahead
<img width="214" height="400" alt="home" src="https://user-images.githubusercontent.com/54924480/208783443-a4dacafa-7c30-430b-b00a-3581eefa8145.png">
Users can see their deadlines today. The application will display the data such as task name, course name, class name, and deadline time. They also can see deadlines seven days ahead.

### 4. View courses, classes, and tasks
<img width="214" height="400" alt="view courses" src="https://user-images.githubusercontent.com/54924480/208783095-83663c85-e7b8-43cb-9dce-8b003f8012e4.png"> <img width="214" height="400" alt="view list class" src="https://user-images.githubusercontent.com/54924480/208783170-c2e3ce82-b9d3-4217-9016-214a7169a035.png"> <img width="214" height="400" alt="view class detail" src="https://user-images.githubusercontent.com/54924480/208783323-c9a8a9a0-6b2d-4b55-88ee-f1fb10164d07.png">

Users can see all available courses, see the classes, and the tasks.

### 5. Profile
<img width="214" height="400" alt="profile" src="https://user-images.githubusercontent.com/54924480/208784173-76f1a32c-48c0-42d6-b1d4-fd3dd4abd72e.png">
Users can see their profile data.

### 6. Change language
<img width="214" height="400" alt="change language" src="https://user-images.githubusercontent.com/54924480/208784312-935fee68-40ae-491c-a016-255f2a69cc73.png">
Users can change the language of the application. Currently available in English and Bahasa Indonesia.

### 7. Admin: CRUD courses and tasks
<img width="214" height="400" alt="admin list courses" src="https://user-images.githubusercontent.com/54924480/208784517-1ca6465a-53ab-4f07-8a9c-448f385b53ee.png"> <img width="214" height="400" alt="add course" src="https://user-images.githubusercontent.com/54924480/208784553-9f58f136-7d8d-4907-bc90-c75fa11c34be.png"> <img width="214" height="400" alt="admin task list" src="https://user-images.githubusercontent.com/54924480/208784622-1a82c81b-e121-44d6-8d84-0a1f62589100.png"> <img width="214" height="400" alt="add task" src="https://user-images.githubusercontent.com/54924480/208784693-3b7d08aa-cff0-4b33-8945-b14e5999eb8e.png"> <img width="214" height="400" alt="select date" src="https://user-images.githubusercontent.com/54924480/208784898-2418318d-9ec3-49a3-bfd8-b0de1bdc4139.png"> <img width="214" height="400" alt="select time" src="https://user-images.githubusercontent.com/54924480/208784916-4bc20267-cfbc-4b0d-86a5-f8eeefdfe9b4.png">



Users granted an admin feature can upload, update, and delete courses and tasks.
