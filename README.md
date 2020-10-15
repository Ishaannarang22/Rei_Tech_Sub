# ReiTechSub

Project Map
1. / or /home (login or signup)(dashboard, if cookie detected)
2. /login (login page) links to :-
3. /setup (setup the device)
4. /signup (signup page)
5. /u/{}/dashboard/master (the users' master dashboard) *
6. /u/{}/dashboard/drone (the users' drone dashboard) *
7. /static/\* (for static files)

**\* Authentication Required**

## Bold Footnotes
1. *The Logins unlike the signups are handled by two different functions because the Hyper Text Transfer Protocol was not made to do the kinds of thing we do with it today and thus, doesn't allow to us to write cookies after a header has been written and for some reason golang fails this silently*
