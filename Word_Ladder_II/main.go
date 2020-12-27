t (
	    transformsAllowed int = 1
		)

		func minConversions(startWord string, word string) bool {
			    count := 0

					    for idx, el := range startWord {
								        if el != rune(word[idx]) {
													            count += 1
																			        }
																							        if count > transformsAllowed {
																												            return false
																																		        }
																																						    }
																																								    return true
																																									}

																																									func findPossiblePerms(endWord string, wordList []string) <-chan string {    
																																										    chnl := make(chan string)
																																												    go func() {
																																															        for _, word := range wordList {
																																																				            if  endWord != word && minConversions(word, endWord)  {
																																																											                chnl <- word
																																																																			            }
																																																																									        }
																																																																													        defer close(chnl)
																																																																																	    }()
																																																																																			    return chnl
																																																																																				}


																																																																																				func findLadders(beginWord string, endWord string, wordList []string) [][]string {
																																																																																					    
																																																																																					    words := map[string]bool{}
																																																																																							    for _, w := range wordList {
																																																																																										        words[w] = true 
																																																																																														    }
																																																																																																    
																																																																																																    if _, ok := words[endWord]; !ok {
																																																																																																			        return [][]string{}
																																																																																																							    }
																																																																																																									    
																																																																																																									    
																																																																																																									    result := [][]string{}
																																																																																																											    level :=  1
																																																																																																													    length := math.MaxInt32
																																																																																																															    visited := map[string]bool{}
																																																																																																																	    queue := [][]string{}
																																																																																																																			    queue = append(queue, []string{beginWord})
																																																																																																																					    
																																																																																																																					    
																																																																																																																					    for len(queue) > 0 {
																																																																																																																								        ladder := queue[0]
																																																																																																																												        queue = queue[1:]
																																																																																																																																        
																																																																																																																																        if len(ladder) > level {
																																																																																																																																					            for w, _ := range visited {
																																																																																																																																												                delete(words, w)
																																																																																																																																																				            }
																																																																																																																																																										            visited = map[string]bool{}
																																																																																																																																																																        }
																																																																																																																																																																				        lastWord := ladder[len(ladder)-1]
																																																																																																																																																																								        if len(ladder) > length {
																																																																																																																																																																													            break
																																																																																																																																																																																			        } else {
																																																																																																																																																																																								            level = len(ladder)
																																																																																																																																																																																														            if lastWord == endWord {
																																																																																																																																																																																																					                result = append(result, ladder)
																																																																																																																																																																																																													                length = len(ladder)
																																																																																																																																																																																																																					            }
																																																																																																																																																																																																																											        }
																																																																																																																																																																																																																															        
																																																																																																																																																																																																																															        for newWord := range findPossiblePerms(lastWord, wordList) {
																																																																																																																																																																																																																																				            if _, ok := words[newWord]; !ok || newWord == lastWord {
																																																																																																																																																																																																																																											               continue 
																																																																																																																																																																																																																																																		             }
																																																																																																																																																																																																																																																								             visited[newWord] = true
																																																																																																																																																																																																																																																														             tempPath := []string{}
																																																																																																																																																																																																																																																																				             tempPath = append(tempPath, ladder...)
																																																																																																																																																																																																																																																																										             tempPath = append(tempPath, newWord)
																																																																																																																																																																																																																																																																																             queue = append(queue, tempPath)
																																																																																																																																																																																																																																																																																						         }
																																																																																																																																																																																																																																																																																										     }
																																																																																																																																																																																																																																																																																												     return result 
																																																																																																																																																																																																																																																																																													 }
