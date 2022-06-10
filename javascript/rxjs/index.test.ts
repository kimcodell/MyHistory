import {
  bufferCount,
  concatMap,
  from,
  mergeMap,
  range,
  Subject,
  takeUntil,
} from "rxjs";
import greet from "../lib/greet"; //import 위해서 babel 세팅해줌

jest.setTimeout(90000);

describe("rxjs_test", () => {
  it("test greet", () => {
    greet();
  });

  it("test_array", (done) => {
    const stocks$ = range(1, 1000); // 주식 코드 1000개로 만든 observable이라고 가정.

    async function buy(stockCode: number) {
      return new Promise((resolve, reject) => {
        setTimeout(() => {
          console.log(`I bought ${stockCode}`);
          resolve(stockCode);
        }, 1000);
      });
    }

    function buyMany$(stockCodes: number[]) {
      return from(stockCodes).pipe(
        mergeMap(async (code) => {
          return await buy(code);
        })
      );
    }

    const stopSignal$ = new Subject();

    setTimeout(() => {
      stopSignal$.next("stop!");
    }, 5000);

    stocks$
      .pipe(
        bufferCount(100),
        concatMap((stockCodes) => buyMany$(stockCodes)),
        takeUntil(stopSignal$)
      )
      .subscribe({
        next: (v) => console.log(v),
        error: (err) => console.error(err),
        complete: () => {
          console.log("complete!");
          done();
        },
      });
  });

  it("mergeMap", () => {});

  it("concatMap", () => {});

  it("subject", () => {});

  it("takeUntil", () => {});

  it("토큰 대량 전송 기능 테스트", () => {
    
  });
});
