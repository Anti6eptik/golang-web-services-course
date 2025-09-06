package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	var depth int
	var printfiles bool
	fmt.Println("Введите название папки")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("Некорректное название папки")
		return
	}
	fmt.Println("Введите максимальную глубину")
	_, err = fmt.Scanln(&depth)
	if err != nil {
		fmt.Println("Некорректная глубина")
		return
	}
	fmt.Println("Введите 0 если хотите визулизировать только папки или 1 если хотите визуализировать папки и файлы")
	_, err = fmt.Scan(&printfiles)
	if err != nil {
		fmt.Println("Некорректный ответ")
		return
	}
	OpenDirs(name, 0, depth, false, "	", printfiles)

}

func OpenDirs(name string, depth int, targetDepth int, last bool, prefix string, printfiles bool) (err error) {
	var flag bool

	dir1, err := os.Open(name)
	if err != nil {
		return err
	}
	defer dir1.Close()
	info1, err := dir1.Stat()
	if err != nil {
		return err
	}
	if depth > targetDepth {
		return nil
	}
	if info1.IsDir() {
		if depth == 0 {
			fmt.Printf("%s\n", info1.Name())

		} else if last {
			fmt.Printf("%s└───%s\n", prefix, info1.Name())
			prefix = prefix + "\t"
		} else {
			fmt.Printf("%s├───%s\n", prefix, info1.Name())
			prefix = prefix + "│\t"
		}

		names1, err := dir1.Readdirnames(-1)
		if err != nil {
			return err
		}
		for i, val1 := range names1 {
			flag = true
			if !printfiles {
				for _, val2 := range names1[i+1:] {
					file, err := os.Open(name + "/" + val2)
					if err != nil {
						return err
					}
					defer file.Close()
					info2, err := file.Stat()
					if err != nil {
						return err
					}
					if info2.IsDir() {
						flag = false
						break
					}
				}
			} else if i != len(names1)-1 {
				flag = false
			}
			if OpenDirs(name+"/"+val1, depth+1, targetDepth, flag, prefix, printfiles) != nil {
				return err
			}
		}
	} else if printfiles {
		if last {
			if info1.Size() == 0 {
				fmt.Printf("%s└───%s (empty)\n", prefix, info1.Name())
			} else {
				fmt.Printf("%s└───%s (%db)\n", prefix, info1.Name(), info1.Size())
			}
		} else {
			if info1.Size() == 0 {
				fmt.Printf("%s├───%s (empty)\n", prefix, info1.Name())
			} else {
				fmt.Printf("%s├───%s (%db)\n", prefix, info1.Name(), info1.Size())
			}
		}

	}
	return nil
}
