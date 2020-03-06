# Xml Tag Check Duplicates

## Usage

```golang
ad := &Ad{}

files := []string{
		"file1.xml",
		"file2.xml",
    		"etc...",
		}
		
	var urls, _ []string = ad.filesProcess(files, ITEM_CONTAINER_URL_TAG_NAME) //example: <ad> "ad"

	//Function to check duplicates
	ad.tagCompare(urls)
```

## Output

Ads processed 10 

https://www.EXAMPLE.com/ITEM/8415/3 Duplicated 3 Times<br/>
https://www.EXAMPLE.com/ITEM/87565/8 Duplicated 2 Times<br/>
https://www.EXAMPLE.com/ITEM/5657/5 Duplicated 3 Times<br/>

Total duplicates 8

2020/03/06 12:51:36 Exec Time 0.07s

## License
[MIT](https://choosealicense.com/licenses/mit/)
