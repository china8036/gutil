package str

import (
	"fmt"
	"strings"
)


func NewSeg()*Seg{
	return &Seg{}
}

type SegDic map[string][]string

type  Seg struct {
	dic SegDic

}

//获取分词结果
func (s *Seg) GetSegDic(title string)SegDic{
	tmp_dic := s.seg(title)
	for index,actor := range tmp_dic{
             s.pushActor(actor,index)
	}
	return s.dic
}


//分词
func (s *Seg)seg(title string) []string {
	tmp_rune := []rune(title)
	tmp_return := make([]string, len(tmp_rune))
	tmp_index := 0
	for _, actor := range tmp_rune {
		ascii_code := int(actor)
		if ascii_code <= 0x9fa5 && ascii_code >= 0x4e00 { //中文一个一个的分 其他见空格分
			if tmp_return[tmp_index] != "" {
				tmp_index++
			}
			tmp_return[tmp_index] = string(actor)
			tmp_index++
			continue
		} else if ascii_code == 0x20 {
			if tmp_return[tmp_index] != "" { //即上一个不是空 如果是空的话 直接覆盖就ok
				tmp_index++
			}
			continue
		}else if ascii_code>= 65 && ascii_code <= 90{//按小写字母存储 以便于查询
			tmp_return[tmp_index] = fmt.Sprintf("%s%s", tmp_return[tmp_index], strings.ToLower(string(actor)))
		}else{
			tmp_return[tmp_index] = fmt.Sprintf("%s%s", tmp_return[tmp_index], string(actor))
		}


	}
	rang := len(tmp_rune)
	if tmp_index+1 < rang {
		rang = tmp_index+1
	}
	return tmp_return[0:rang]
}


//过滤重复和记录位置
func (s *Seg) pushActor(str string,index int) {
	if s.dic == nil{
		s.dic = make(SegDic)
	}
        if s.dic[str] == nil{
		s.dic[str] = make([]string,0)
	}
	s.dic[str] = append(s.dic[str],fmt.Sprintf("%d",index))

}