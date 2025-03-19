package randresources

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Image struct {
	value  string
	img    *image.RGBA
	buffer *bytes.Buffer
}

func getColorFromHash(hash string) color.RGBA {
	hexColor := hash[:6]
	r, _ := strconv.ParseUint(hexColor[0:2], 16, 8)
	g, _ := strconv.ParseUint(hexColor[2:4], 16, 8)
	b, _ := strconv.ParseUint(hexColor[4:6], 16, 8)
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}
}

func generateIdenticon(input string) *image.RGBA {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	// 图案尺寸（块数）
	gridSize := 5
	blockSize := 40 // 每个块的像素大小
	imgSize := gridSize * blockSize
	img := image.NewRGBA(image.Rect(0, 0, imgSize, imgSize))

	// 背景颜色（浅灰色）
	backgroundColor := color.RGBA{R: 240, G: 240, B: 240, A: 255}

	// 主要颜色
	mainColor := getColorFromHash(hashString)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 使用哈希值决定哪些块被填充
	r.Seed(time.Now().UnixNano() + int64(hashString[0])) // 使用哈希值的第一个字符作为种子

	pattern := make([][]bool, gridSize)
	for i := 0; i < gridSize; i++ {
		pattern[i] = make([]bool, gridSize)
	}

	// 生成对称的图案
	for i := 0; i < gridSize; i++ {
		for j := 0; j < (gridSize+1)/2; j++ { // 只需遍历一半的列
			// 使用哈希值的某一位来决定是否填充
			index := i*gridSize + j
			if index < len(hashString) {
				bit := hashString[index]
				if bit >= '8' { // 可以根据需要调整阈值
					pattern[i][j] = true
					pattern[i][gridSize-1-j] = true // 实现水平对称
				}
			} else {
				// 如果哈希值不够长，可以使用随机性
				if rand.Intn(2) == 0 {
					pattern[i][j] = true
					pattern[i][gridSize-1-j] = true
				}
			}
		}
	}

	// 填充像素
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			blockColor := backgroundColor
			if pattern[i][j] {
				blockColor = mainColor
			}
			// 填充当前块的像素
			for y := i * blockSize; y < (i+1)*blockSize; y++ {
				for x := j * blockSize; x < (j+1)*blockSize; x++ {
					img.Set(x, y, blockColor)
				}
			}
		}
	}

	return img
}

func (img *Image) Base64() string {
	if img.buffer == nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(img.Bytes())
}
func (img *Image) Bytes() []byte {
	if img.buffer == nil {
		return nil
	}
	return img.buffer.Bytes()
}

func (img *Image) Save(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img.img)
	if err != nil {
		return err
	}
	return nil
}
func BuildImage(value string) (Image, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		value = createUUID()
	}

	img := generateIdenticon(value)
	i := Image{value, img, nil}

	i.buffer = new(bytes.Buffer)
	err := png.Encode(i.buffer, img)
	if err != nil {
		fmt.Print(err)
		return Image{}, err
	}

	return i, nil
}
