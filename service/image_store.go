package service

import(
	"sync"
	"bytes"
	"github.com/google/uuid"
	"os"
	"fmt"
)


type ImageStore interface{
	Save(laptopID string, imageType string, imageData bytes.Buffer) (string,error)
}

type DiskImageStore struct{
	mutex sync.RWMutex
	imageFolder string
	images map[string]*ImageInfo 
}

type ImageInfo struct{
	LaptopID string
	Type string
	Path string
}

func NewDiskImageStore(path string) *DiskImageStore{
	return &DiskImageStore{
		imageFolder : path,
		images : make(map[string]*ImageInfo),
	}
}

func(store *DiskImageStore) Save(laptopID string, imageType string, imageData bytes.Buffer) (string,error){
	imageID,err := uuid.NewRandom()
	if err != nil{
		return "",err
	}

	imagePath := fmt.Sprintf(store.imageFolder,imageID, imageType)

	file,err := os.Create(imagePath)
	if err != nil{
		return "",nil
	}

	_,err = imageData.WriteTo(file)
	if err != nil{
		return "",nil
	}

	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.images[imageID.String()] = &ImageInfo{
		LaptopID : laptopID,
		Type : imageType,
		Path : imagePath,
	}
	return imageID.String(),nil
}