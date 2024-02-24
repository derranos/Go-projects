import sys
import pandas as pd
import json

df1 = pd.read_json(sys.argv[1])
df2 = pd.read_json(sys.argv[2])
if len(df1.iloc[0]) != len(df2):
    print("Invalid range of matrix:", len(df1.iloc[0]), "!=", len(df2))
    sys.exit()
ans = (df1 @ df2).values.tolist()
print("Answer:")
for row in ans:
    for val in row:
        print(val, end = " ")
    print()