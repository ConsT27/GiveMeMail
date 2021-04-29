package deal


//the func is used to Dereplication
func Dereplication(mailreusltraw []string) []string{
	result := []string{}
	tempMap := map[string]byte{}
	for _, e := range mailreusltraw{
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l{
			result = append(result, e)
		}
	}
	return result
}
