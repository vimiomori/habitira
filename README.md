# habitira

## Habitica-Jira integration

## User setup
  - CLI (habitira)
  - saved in mongodb on heroku deployed API
    - jira token
    - habitica token
    - habitica user ID
    - jira status categories 
    - jira projects
    - tag to give habitica issues from jira
    - story point field?
    - story point to difficulty translation?
    - user rules:
      - on task completion update jira issue with what status category?
      - on task delete update jira issue with what status category?

## Syncing Jira to Habitica
  - cronjob hosted on heroku?
  - cronjob fetches all jira issues matching filters configured by the user
    and calls heroku API to upsert fetched issues. If there are newly created
    issues, cronjob calls habitica API to create the new issues with 
    the user defined tag, and optionally prefix titles an image or an emoji.
  
## Syncing Habitica to Jira
  - cloud function hosted on GCP as webhook target
  - cloud function calls heroku hosted API with user ID and maybe token
