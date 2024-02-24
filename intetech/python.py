import sys
import pandas as pd
import json

array1_data = json.loads(sys.argv[1])
array2_data = json.loads(sys.argv[2])
if len(array1_data[0]) != len(array2_data):
    print("Invalid range of matrix")
    sys.exit()
df1 = pd.DataFrame(array1_data)
df2 = pd.DataFrame(array2_data)
ans = (df1 @ df2).values.tolist()
print("Answer:")
for row in ans:
    for val in row:
        print(val, end = " ")
    print()
