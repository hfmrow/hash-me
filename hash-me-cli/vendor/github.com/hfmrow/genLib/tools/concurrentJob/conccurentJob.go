// concurrentJob.go

/*
	Copyright (c) 2021 H.F.M
	See the The MIT License (MIT) for details:
	https://opensource.org/licenses/mit-license.php

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	SOFTWARE.

	- The purpose of this library is to provide a convenient method to handle
	  concurrent processing.

	ex: where 'files' is []string, 'item' is filename transtyped to .(string)

			ccs = ConcurrentCalcStrucNew(runtime.NumCPU(),

				func(item interface{}) {
					if result, err = makeHash(item.(string)); err == nil && len(result.outStr) > 0 {
						result.content = strings.Join(result.outStr, GetOsLineEnd())
						resultsHash = append(resultsHash, result)
					}
					Logger.Log(err, "doIt/makeHash")
					glib.IdleAdd(func() {
						mainStatusbar.Set(fmt.Sprintf("%v", ccs.Count), 1)
				})

			})
			ccs.Run(ccs.StringSliceToIfaceSlice(files))
			ccs.Wait()
*/

package concurrent

import "sync"

// ConcurrentCalcStruc: structure to hold information and
// callback function to use while doing concurrent operations
type ConcurrentCalcStruc struct {
	CpuCount,
	itemCount,
	cpuNb,
	Count,
	toAdd int

	wgIn,
	wgOut *sync.WaitGroup

	Callback,
	localCallback func(item interface{})
}

// ConcurrentCalcStrucNew: Create new structure
func ConcurrentCalcStrucNew(cpuCount int, callBack func(item interface{})) *ConcurrentCalcStruc {
	ccs := new(ConcurrentCalcStruc)
	ccs.Callback = callBack
	ccs.CpuCount = cpuCount
	ccs.wgIn = new(sync.WaitGroup)
	ccs.wgOut = new(sync.WaitGroup)
	ccs.localCallback = func(item interface{}) {
		ccs.Callback(item)
		ccs.wgOut.Done()
		ccs.wgIn.Done()
		ccs.Count++
	}
	return ccs
}

// StringSliceToIfaceSlice: Convert string slice to interface slice
func (ccs *ConcurrentCalcStruc) StringSliceToIfaceSlice(inStr []string) []interface{} {
	out := make([]interface{}, len(inStr))
	for idx, item := range inStr {
		out[idx] = item
	}
	return out
}

// Wait: Until operations done
func (ccs *ConcurrentCalcStruc) Wait() {
	ccs.wgOut.Wait()
	ccs.Count = ccs.itemCount
}

// Run: Start concurrent computing operations
func (ccs *ConcurrentCalcStruc) Run(items []interface{}) {
	ccs.cpuNb = 1
	ccs.itemCount = len(items)
	ccs.Count = 0

	ccs.wgOut.Add(ccs.itemCount)
	ccs.wgIn.Add(ccs.CpuCount)
	for _, item := range items {

		go ccs.localCallback(item)

		if ccs.cpuNb == ccs.CpuCount {
			ccs.wgIn.Wait()
			ccs.toAdd = ccs.CpuCount
			if (ccs.itemCount - ccs.Count) < ccs.CpuCount {
				ccs.toAdd = ccs.itemCount - ccs.Count
			}
			ccs.wgIn.Add(ccs.toAdd)
			ccs.cpuNb = 0
		}
		ccs.cpuNb++
	}
}
