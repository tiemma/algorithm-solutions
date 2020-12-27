
def get_distinct_count(a = 100, b = 100):
  mem = set()
  count = 2
  while a >= count:
    temp = 2
    while b >= temp:
      mem.add(count**temp) 
      temp += 1
    count+=1
  print(mem)
  return len(set(mem))


if __name__ == "__main__":
  print(get_distinct_count(100, 100))

