import json


'''
DB = {}

sessinID = '1111'
UPF_address = '10.0.0.1'
UE_address = '192.168.11.1'
DB = {sessinID:{'UPF':UPF_address,'UE':UE_address}}

print(DB)
'''

def main():
  DB = {}
  DB = jsonRead()
  #jsonRead()

  #print(DB)
  sessinID = '1112'
  UPF_address = '10.0.0.1'
  UE_address = '192.168.12.1'
  DB[sessinID] = {'UPF':UPF_address,'UE':UE_address}
  #print(DB)

  jsonWrite(DB)

  DB.pop('1111')
  #print(DB)

  jsonWrite(DB)


def jsonRead():
  #data = open('data.json', 'r')
  #data = json.load(data)
  with open('data.json','r') as f:
    data = json.load(f)
  
  #print(data)
  return data

def jsonWrite(data):
  with open('data.json', 'w') as f:
    json.dump(data, f, indent=2)


if __name__ == "__main__":
  #jsonRead()
  main()
