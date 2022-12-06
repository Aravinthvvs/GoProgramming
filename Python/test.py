def findNum(nums,key):
    for i in range(len(nums)):
        if key == nums[i]:
            return True
    return False

if __name__ == "__main__":
    k = int(input())
    nums = [1,2,3,4,5,6]
    if findNum(nums,k) == True:
        print(" Number {} is found in the list".format(k))
    else:
        print(f"Number {k} not found in the list")