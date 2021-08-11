from rich import print
personList = []
for i in range(1,10,1):
    personList.append({
        "id":i,
        "name":i,
        "comment":f"{i}写的评论",
        "read":False
    })

while True:
    print("select num\n1.look all comments\n2.select a comment")
    admininput = input("")
    if admininput == "1":
       print(personList)
    elif admininput == "2":
        print("select persons' id")
        currInput = input("")
        for person in personList:
            if person["id"] == int(currInput):
                person["read"]=True
    else:
        print("error")
