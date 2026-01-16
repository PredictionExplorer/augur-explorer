    var hash_title [32]byte
    title_bytes := []byte(rec.Question)
    sha_title := sha3.NewLegacyKeccak256()
    if _, err := sha_title.Write(title_bytes[:]); err != nil {
        ss.Log_msg(fmt.Sprintf("Error in keccak256(): %v\n",err))
    } else {
        sha_title.Sum(hash_title[:0])
        hex_hash_title := hex.EncodeToString(hash_title[:])
        fmt.Printf("Hash title = %v\n",hex_hash_title)
    }   
    desc_bytes := []byte(rec.Description)
    _=title_bytes
    data := make([]byte,32+len(desc_bytes))
    copy(data[:],hash_title[:])
    copy(data[32:],desc_bytes)
    fmt.Printf("data = %v\n",hex.EncodeToString(data[:]))
    sha := sha3.NewLegacyKeccak256()
    if _, err := sha.Write(data[:]); err != nil {
        ss.Log_msg(fmt.Sprintf("Error in keccak256(): %v\n",err))
    } else {
        var qhash [32]byte
        sha.Sum(qhash[:0])
        hex_qhash := hex.EncodeToString(qhash[:])
        fmt.Printf("Question id = %v\n",rec.QuestionId)
        fmt.Printf("Calculated hash = %v\n",hex_qhash)
    }   
    fmt.Printf("question id = %v\n",rec.QuestionId)
