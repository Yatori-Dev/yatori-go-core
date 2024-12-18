package xuexitong

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// PullChapter ??????????cpi ?key ?? ????
// cpi ? key ?? ????json?????? int
// TODO???? int ???????? ?Course???????? ? ?action?????XueXiTCourseDetailForCourseIdAction???? ???
func (cache *XueXiTUserCache) PullChapter(cpi int, key int) (string, error) {
	method := "GET"

	params := url.Values{}
	params.Add("id", strconv.Itoa(key))
	params.Add("personid", strconv.Itoa(cpi))
	params.Add("fields", "id,bbsid,classscore,isstart,allowdownload,chatid,name,state,isfiled,visiblescore,hideclazz,begindate,forbidintoclazz,coursesetting.fields(id,courseid,hiddencoursecover,coursefacecheck),course.fields(id,belongschoolid,name,infocontent,objectid,app,bulletformat,mappingcourseid,imageurl,teacherfactor,jobcount,knowledge.fields(id,name,indexOrder,parentnodeid,status,isReview,layer,label,jobcount,begintime,endtime,attachment.fields(id,type,objectid,extension).type(video)))")
	params.Add("view", "json")

	client := &http.Client{}
	req, err := http.NewRequest(method, ApiPullChapter+"?"+params.Encode(), nil)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("User-Agent", " Dalvik/2.1.0 (Linux; U; Android 12; SM-N9006 Build/70e2a6b.1) (schild:e9b05c3f9fb49fef2f516e86ac3c4ff1) (device:SM-N9006) Language/zh_CN com.chaoxing.mobile/ChaoXingStudy_3_6.3.7_android_phone_10822_249 (@Kalimdor)_4627cad9c4b6415cba5dc6cac39e6c96")
	req.Header.Add("Accept-Language", " zh_CN")
	req.Header.Add("Host", " mooc1-api.chaoxing.com")
	req.Header.Add("Connection", " Keep-Alive")
	req.Header.Add("Cookie", cache.cookie)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(body), nil
}

// FetchChapterPointStatus 章节状态
// nodes 各章节对应ID
func (cache *XueXiTUserCache) FetchChapterPointStatus(nodes []int, clazzID, userID, cpi, courseID int) (string, error) {
	method := "POST"
	strInts := make([]string, len(nodes))
	for i, v := range nodes {
		strInts[i] = fmt.Sprintf("%d", v)
	}

	ts := time.Now().UnixNano() / 1000000
	join := strings.Join(strInts, ",")
	values := url.Values{
		"view":     {"json"},
		"nodes":    {join},
		"clazzid":  {strconv.Itoa(clazzID)},
		"time":     {strconv.FormatInt(ts, 10)},
		"userid":   {strconv.Itoa(userID)},
		"cpi":      {strconv.Itoa(cpi)},
		"courseid": {strconv.Itoa(courseID)},
	}
	// 编码请求体
	payload := strings.NewReader(values.Encode())

	client := &http.Client{}
	req, err := http.NewRequest(method, ApiChapterPoint, payload)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "mooc1-api.chaoxing.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", cache.cookie)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	// 解码响应体（假设服务器返回的内容是 ISO-8859-1 编码）
	// decodedBody, _, err := transform.Bytes(charmap.ISO8859_1.NewDecoder(), body)
	return string(body), nil
}

// FetchChapterCords 以课程序号拉取对应“章节”的任务节点卡片资源
// Args:
//
//	nodes: 任务点集合 , index: 任务点索引
func (cache *XueXiTUserCache) FetchChapterCords(nodes []int, index, courseId int) (string, error) {
	method := "GET"
	values := url.Values{}
	values.Add("id", strconv.Itoa(nodes[index]))
	values.Add("courseid", strconv.Itoa(courseId))
	values.Add("fields", "id,parentnodeid,indexorder,label,layer,name,begintime,createtime,lastmodifytime,status,jobUnfinishedCount,clickcount,openlock,card.fields(id,knowledgeid,title,knowledgeTitile,description,cardorder).contentcard(all)")
	values.Add("view", "json")
	values.Add("token", "4faa8662c59590c6f43ae9fe5b002b42")

	client := &http.Client{}
	req, err := http.NewRequest(method, ApiChapterCards+"?"+values.Encode(), nil)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("User-Agent", " Dalvik/2.1.0 (Linux; U; Android 12; SM-N9006 Build/70e2a6b.1) (schild:e9b05c3f9fb49fef2f516e86ac3c4ff1) (device:SM-N9006) Language/zh_CN com.chaoxing.mobile/ChaoXingStudy_3_6.3.7_android_phone_10822_249 (@Kalimdor)_4627cad9c4b6415cba5dc6cac39e6c96")
	req.Header.Add("Accept-Language", " zh_CN")
	req.Header.Add("Host", " mooc1-api.chaoxing.com")
	req.Header.Add("Connection", " Keep-Alive")
	req.Header.Add("Cookie", cache.cookie)
	req.Header.Add("Accept", "*/*")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	return string(body), nil
}
