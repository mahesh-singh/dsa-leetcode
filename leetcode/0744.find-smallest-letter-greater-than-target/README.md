# Sudo code


```
l=0
r=len(letters)-1
ret = letters[0]
while l < r
    mid= l+ (r-l)/2

    if letters[mid] == target 
        return letters[mid+1]
    
    if letters[mid] > target
        //move right
        ret = letters[mid]
        r = mid-1
    else
        l = mid +1
    
return ret
     

```