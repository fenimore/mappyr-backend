#from nose.tools import assert_true
import requests

# Example tokens with id

token_1 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyLWlkIjoiMSIsImV4cCI6MTQ3Njk3MTQ2NywiaXNzIjoibG9jYWxob3N0OjgwODAifQ.wv1UgclK5uKUYYZFpnx4DjLcHlinirTfzL0nhmrJ7gc"

token_8 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyLWlkIjoiOCIsImV4cCI6MTQ3Njk3MzEwNCwiaXNzIjoibG9jYWxob3N0OjgwODAifQ.NnxhEH7ETklp8H_hkNNxFaJpHt0s4TdE2gWVRpSE39Q"

##########################################################################################
#USER ACTIONS
##########################################################################################
def signup():
    """Create a new user"""
    r = requests.post('http://localhost:8080/signup', json ={"username":"Wrinkle",
                                                             "password":"Woootwooot",
                                                             "email":"Noootme@gmail.com"})
    print(r.status_code)
    print(r.json())


def login():
    """Login a user with a username and password"""
    r = requests.post('http://localhost:8080/login', json ={"username":"Simone",
                                                             "password":"lulz"})
    print(r.status_code)
    print(r.json())

def comment():
    """Post a comment with a token in the headers"""
    r = requests.post('http://localhost:8080/new',
                      json={"title":"SUsh1","description":"THIS Rox0rs",
                            "latitude":45, "longitude":88},
                      headers={"Authentication":token_8})

    print(r.status_code)
    print(r.json())


def upvote(comment):
    """UpVOTE!"""
    r = requests.get('http://localhost:8080/upvote/'+ comment, headers={"Authentication":token_8})
    print(r.status_code)
    print(r.json())

def downvote(comment):
    """downVOTE!"""
    r = requests.get('http://localhost:8080/downvote/'+ comment, headers={"Authentication":token_8})
    print(r.status_code)
    print(r.json())

##########################################################################################
#Get Votes or COmments of User
##########################################################################################
def user_votes(id):
    """User votes takes the id, but pass a string"""
    r = requests.get("http://localhost:8080/votes/"+id)

    print(r.status_code)
    print(r.json())

def user_comments(id):
    """User comments takes the id, but pass a string"""
    r = requests.get("http://localhost:8080/comments/"+id)

    print(r.status_code)
    print(r.json())
