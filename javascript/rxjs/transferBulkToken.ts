import { bufferCount, from } from "rxjs";

const transferBulkToken = () => {
  const data = [
    { address: "0x4dd3840597c00ec2773043b2a6a98e1c45aeddd6", tokenId: 22},
    { address: "0x4dd3840597c00ec2773043b2a6a98e1c45aeddd6", tokenId: 22},
    { address: "0x4dd3840597c00ec2773043b2a6a98e1c45aeddd6", tokenId: 22},
    { address: "0x4dd3840597c00ec2773043b2a6a98e1c45aeddd6", tokenId: 22},
    { address: "0x4dd3840597c00ec2773043b2a6a98e1c45aeddd6", tokenId: 22},
    { address: "0x4dd3840597c00ec2773043b2a6a98e1c45aeddd6", tokenId: 22},
    { address: "0x4dd3840597c00ec2773043b2a6a98e1c45aeddd6", tokenId: 22},
    { address: "0x4dd3840597c00ec2773043b2a6a98e1c45aeddd6", tokenId: 22},
    { address: "0x4dd3840597c00ec2773043b2a6a98e1c45aeddd6", tokenId: 22},
    { address: "0x4dd3840597c00ec2773043b2a6a98e1c45aeddd6", tokenId: 22},
    { address: "0x4dd3840597c00ec2773043b2a6a98e1c45aeddd6", tokenId: 22},
    { address: "0x4dd3840597c00ec2773043b2a6a98e1c45aeddd6", tokenId: 22},
    { address: "0x4dd3840597c00ec2773043b2a6a98e1c45aeddd6", tokenId: 22},
    { address: "0x4dd3840597c00ec2773043b2a6a98e1c45aeddd6", tokenId: 22},
  ]
  from(data).pipe(
    bufferCount(10),

  );
};

export default transferBulkToken;
