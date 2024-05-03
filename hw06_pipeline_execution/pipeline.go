package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in

	for _, stage := range stages {
		stageResult := make(Bi)

		go func(resultCh Bi, out Out) {
			defer close(resultCh)

			for {
				select {
				case <-done:
					return
				case v, ok := <-out:
					if !ok {
						return
					}
					resultCh <- v
				}
			}
		}(stageResult, out)

		out = stage(stageResult)
	}

	return out
}
