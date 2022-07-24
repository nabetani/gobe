import glob
from tokenize import group
import matplotlib.pyplot as plt
import numpy as np
import matplotlib
import os.path
import re

matplotlib.use('Agg')

def readall(path):
    with open(path) as f:
        return f.readlines()

def readdata(path):
    data = readall(path)
    env = None
    lang = None
    r={}
    for line in data:
        if line.startswith("## "):
            env = line[3:].strip()
        elif line.startswith("### "):
            lang = line[4:].strip()
        else:
            m=re.search(r"8223390\,\s*(.*), tick.*?(\d(?:.*))ms", line)
            if m:
                c,t = m.group(1), m.group(2)
                if not r.get(env): r[env]={}
                if not r[env].get(lang): r[env][lang]={}
                if not r[env][lang].get(c): r[env][lang][c]=[]
                r[env][lang][c].append(float(t))
    return r


def rename(x):
    m=re.search(r"(go\d+\.\d+)", x)
    if m:
        return m.group(1)
    m=re.search(r"clang", x)
    if m:
        return "clang"
    m=re.search(r"11.3.0", x)
    if m:
        return "g++-11"

def main():
    a,_ = os.path.split(__file__)
    data = readdata(a+"/../README.md")
    print(repr(data))
    fig = plt.figure(figsize=(6, 3), dpi=200)
    envs = ['intel MBP', 'M1 Pro 非MAX MBP']
    for i in range(len(envs)):
        env = envs[i]
        ax = fig.add_subplot(1, 2, i+1)
        goLabels = [x for x in data[env]["go"]]
        cppLabels = [x for x in data[env]["c++"]]
        goTicks = [min(data[env]["go"][x]) for x in goLabels]
        cppTicks = [min(data[env]["c++"][x]) for x in cppLabels]
        ticks = goTicks + cppTicks
        labels = [rename(x) for x in goLabels + cppLabels]
        left = list(range(len(ticks)))
        ax.barh( left, ticks,  tick_label=labels, align="center")
    fig.subplots_adjust()
    plt.tight_layout()
    plt.savefig("graph.png")

main()