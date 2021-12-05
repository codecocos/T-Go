package dataStruct

func Hash(s string) int {
	h := 0
	A := 256
	B := 3571 //해쉬값은 0 ~ 3570

	for i := 0; i < len(s); i++ {
		h = (h*A + int(s[i])) % B
	}
	return h
}

type keyValue struct {
	key   string
	value string
}

//해쉬 값의 범위에 해당하는 배열 생성
type Map struct {
	keyArray [3571][]keyValue
}

//빈 맵을 만들어서 반환
func CreateMap() *Map {
	return &Map{}
}

func (m *Map) Add(key, value string) {
	//키를 해쉬로 돌려서 나온 값
	h := Hash(key)
	//키를 해쉬로 돌려서 나온 값의 리스트에 키밸류를 추가후 재정의
	//해쉬의 인덱스에 해당하는 배열의 리스트에다가 키밸류를 추가후 재정의
	m.keyArray[h] = append(m.keyArray[h], keyValue{key, value})
}

func (m *Map) Get(key string) string {
	h := Hash(key)
	for i := 0; i < len(m.keyArray[h]); i++ {
		if m.keyArray[h][i].key == key {
			return m.keyArray[h][i].value
		}
	}
	return ""
}