# Election API

## **⭐** Minimum Viable Product (MVP)

- New user can register account to the system ✔️
- User can login to the system ✔️
- User can edit their profile account ✔️
- User can view the candidate's posts ✔️
- User can comment on candidate’s posts ✔️
- Users can view information about the candidates ✔️
- Users can cast their votes for candidates during the specified election period ✔️
- Admin can promote user to candidate ✔️
- Admin can view the candidate’s posts ✔️
- Admin can set the start and end dates for the election period ✔️
- Admin can delete the user/candidate ✔️
- Admin can delete the candidate's posts ✔️
- Admin can delete user comment ✔️
- Candidate can create, update, delete a post ✔️

## **[🌎](https://emojipedia.org/globe-showing-americas)** Service Implementation

```
GIVEN => I am a new user
WHEN  => I register to the system
THEN  => System will record and return the visitor's username

GIVEN => I am a user
WHEN  => I log in to the system
THEN  => System will authenticate and grant access based on user credentials

GIVEN => I am a user
WHEN  => I edit my profile account
THEN  => The system will update my account with the new details

GIVEN => I am a user
WHEN  => I view a candidate's post
THEN  => System will display the selected candidate's post along with its details

GIVEN => I am a user
WHEN  => I comment on a candidate’s post
THEN  => System will record my comment and return it under the candidate’s post

GIVEN => I am a user
WHEN  => I took an action to view candidate's posts
THEN  => System will show a candidate's post

GIVEN => I am a user
WHEN  => I cast my vote for a candidate during the specified election period
THEN  => System will register my vote for the selected candidate

GIVEN => I am an admin
WHEN  => I promote a user to a candidate
THEN  => System will update the user's status to candidate

GIVEN => I am an admin
WHEN => I view a candidate’s posts
THEN => System will display the posts created by the candidate

GIVEN => I am an admin
WHEN  => I set the start and end dates for the election period
THEN  => System will update the election period accordingly

GIVEN => I am an admin
WHEN  => I delete a user or candidate
THEN  => System will remove the user/candidate from the system

GIVEN => I am an admin
WHEN  => I delete a candidate’s post
THEN  => System will show a deletion status message and delete relevant post

GIVEN => I am an admin
WHEN => I delete a user comment
THEN => System will remove the user comment from the candidate's post

GIVEN => I am a candidate
WHEN  => I create a new post
THEN  => System will record and show the creation status message

GIVEN => I am a candidate
WHEN  => I update my post
THEN  => System will apply the changes and show an update status message

GIVEN => I am a candidate
WHEN  => I delete one of my posts
THEN  => System will show a deletion status message and delete relevant post
```

## **👪** Entities and Actors

```
## Entities ##
User: Individuals who interact with the system. They can have different roles such as regular users, candidates, or administrators.
Attributes:
- id
- username
- password
- name
- bio
- role (user | candidate | admin)
- can_vote
- created_at
- updated_at
- deleted_at

Candidate: Represents users who are running for elections
Attributes:
- id
- user_id
- vote_count
- created_at
- updated_at
- deleted_at

Post: Represents the content created by candidates.
Attributes:
- id
- title
- content
- candidate_id
- created_at
- updated_at
- deleted_at

Comment: Represents the comments made by users on candidates posts
Attributes:
- id
- content
- user_id
- post_id
- created_at
- updated_at
- deleted_at

ElectionPeriod: Represents the period of election event
Attributes include:
- id
- start_date
- end_date
- created_at
- updated_at
- deleted_at
```

```
## Actors ##
- User: Individuals who interact with the system. They can have different roles such as regular users, candidates, or administrators.
- Candidate: Users who are running for election. They can create posts to campaign and receive comments from other users.
- Administrator: Users with special privileges to manage the system, such as election period, users, candidates, posts, and comments.
```

## **🧪** API Installation

### Prerequisites
Before you begin, ensure you have the these installed on your machine:

- [Go](https://go.dev/doc/install)
- [MySQL Server](https://dev.mysql.com/downloads/)

### Running
1. Make sure your MySQL server is configured correctly


2. Clone the repository or download manually:
```shell
git clone --branch i-putu-natha-kusuma https://github.com/nathakusuma/election-api
```


3. Navigate to the project directory:
```shell
cd election-api
```


4. Install dependencies:
```shell
go get .
```


5. Configure your credentials and database configuration in `.env`. For reference, check `.env.example`.


6. Start the application:
```shell
go run main.go
```

## 📃 **Api Documentation**
https://documenter.getpostman.com/view/32594897/2s9YyzeJWk
