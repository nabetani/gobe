import glob
from tokenize import group
import matplotlib.pyplot as plt
import numpy as np
import matplotlib
import os.path
import re

matplotlib.use('Agg')
plt.rcParams['font.family'] = 'Hiragino Sans'

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

BH=0.45

def pos(key,e):
    m=re.search(r"go\d+\.(\d+)", key)
    if m:
        return int(m.group(1))-18 + e*BH
    m=re.search(r"clang", key)
    if m:
        return 1.5 + e*BH
    m=re.search(r"11.3.0", key)
    if m:
        return 2.5 + e*BH


def main():
    a,_ = os.path.split(__file__)
    data = readdata(a+"/../README.md")
    fig = plt.figure(figsize=(6, 3), dpi=200)
    envs = ['intel MBP', 'M1 Pro 非MAX MBP']
    labels={}
    ax = fig.add_subplot(1, 1,1)
    for e in range(len(envs)):
        env = envs[e]
        edata = dict(data[env]["go"], **data[env]["c++"])
        for key in edata:
            val = min(edata[key])
            print(key,val)
            ax.barh( [pos(key,e)], [val], height=BH, align="center", color="br"[e])
            labels[pos(key,0)] = rename(key)

    print(repr((list(labels.keys()), list(labels.values()))))
    ax.set_yticks(list(labels.keys()), list(labels.values()))
    ax.set_xlabel("処理時間(ms)")
    fig.subplots_adjust()
    plt.tight_layout()
    plt.savefig("graph.png")
main()