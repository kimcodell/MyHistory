const greet = () => {
  console.log('greet');

  /*
    X 상품 가격
    1.1X = 서비스 운영에 필요한 금액
      (전송 수수료 + 운영 비용 => 우선 10%로 가정. 추후 정확한 비용 구조 계산 필요);
    C 투표 참여 가격
    P 투표 참여자 수

    1명 -C원 + 상품
    0.5 * P -1 명: 0원
    0.5 * P 명: -C원

    (0.5P + 1) * C = 1.1X

    C = 1.1X / (0.5P + 1);

    ex) 에어팟 20만원, 100명 모집
    (200,000원 * 1.1) / 51명 = 4,500원

    목표 인원을 정해두고 진행할 수도 있고, 일정기간을 정해두고 진행하는 방법도 있음.

    한 사람이 여러 지갑 주소로 참여해도 괜찮음.
  
    <가능한 악의적인 참여 case>
  
    
    <필요한 기능>
    [프론트]
    지갑 연동

    이벤트 리스트 확인

    이벤트 상세

    이벤트 참여

    [백]
    이벤트 목록

    이벤트 참여(일단 취소 불가)

    목표 인원 수 || 정해진 기간 도달 시 당첨자 뽑기 => 바로 정산까지

    [스마트 컨트랙트]
    전송

    당첨자 선정
  */
};

export default greet;
