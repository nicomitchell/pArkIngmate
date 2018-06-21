import subprocess

option = 1

cmd = "sudo rm current.jpg"
subprocess.call(cmd.split(),cwd="../../darknet/data")

cmd_options = ["wget http://31.16.111.66:8083/record/current.jpg","wget http://80.92.117.241/record/current.jpg", "wget http://83.162.139.253/record/current.jpg"]
cmd = cmd_options[option]
print(subprocess.check_output(cmd.split(),cwd="../../darknet/data"))

cmd3 = "./darknet detect cfg/yolo.cfg yolo.weights data/current.jpg -thresh3"
output = str(subprocess.check_output(cmd3.split(), cwd="../../darknet"))
print(output)

separator = "\n\n===============\nLABELS : BEGIN\n===============\n"
try:
    ignore, keep = output.split(separator)
except ValueError:
    print("could not split line", output)

print("\n\n\n" + keep)

items = keep.split("\t\t\n")
remove = False
i = 0
while i < len(items):
    if i % 2 == 0 or i == 0:
        if items[i][:3] != "car":
            print("REMOVED: " + items[i])
            items.remove(items[i])
            if i < len(items)-1:
                items.remove(items[i])          #remove bounding boxes as well            
            continue
    i += 1
# for i, val in  enumerate(items):
#     if remove:
#         print("REMOVED: " + val)
#         items.remove(val)
#         continue
#     remove = False
#     print ("\n" + val)
#     if i % 2 == 0 or i == 0:
#         if val[:3] != "car":
#             print("REMOVED: " + val)
#             items.remove(val)
#             remove = True

for val in items:
    if val[:3] == "car":
        print("REMOVED: " + val)
        items.remove(val)
print("With extras removed:\n")
for i in items:
    print(i + "\n")


f = open("../bounding_box.txt", 'w')

for i in items:
    i = i.replace("\\t","\t")
    print("writing " + i + " to coordinate file")
    f.write(i + "\n")

f.close()
