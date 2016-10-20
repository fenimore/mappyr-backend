#from nose.tools import assert_true
import requests

token_1 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyLWlkIjoiMSIsImV4cCI6MTQ3Njk3MTQ2NywiaXNzIjoibG9jYWxob3N0OjgwODAifQ.wv1UgclK5uKUYYZFpnx4DjLcHlinirTfzL0nhmrJ7gc"

def signup():
    r = requests.post('http://localhost:8080/signup', json ={"username":"simone",
                                                             "password":"frisky",
                                                             "email":"debeav@gmail.com"})
    print(r.status_code)
    print(r.json())
