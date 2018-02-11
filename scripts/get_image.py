import subprocess

cmd = "sudo rm current.jpg"
subprocess.call(cmd.split(),cwd="../../darknet/data")

cmd = "wget http://31.16.111.66:8083/record/current.jpg"
print(subprocess.check_output(cmd.split(),cwd="../../darknet/data"))

cmd3 = "./darknet detect cfg/yolo.cfg yolo.weights data/current.jpg -thresh5"


output = str(subprocess.check_output(cmd3.split(), cwd="../../darknet"))
print(output)

separator = "\\n\\n===============\\nLABELS : BEGIN\\n===============\\n"
ignore,keep = output.split(separator)

print("\n\n\n" + keep)

