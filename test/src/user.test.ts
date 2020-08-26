import proto from "protobufjs";
import path from "path"
import Axios from "axios";

test("should load user.proto", async () => {
  const root = proto.loadSync([
    path.resolve(__dirname, "../../protobuf/caesar/user.proto"),
    path.resolve(__dirname, "../../protobuf/caesar/basic.proto"),
  ]);

  expect(root).toBeTruthy();

  const UserRegistration = root.lookupType("UserRegistration");

  const payload = {
    username: "user",
    password: "pass",
    email: "em",
  };

  let err = UserRegistration.verify(payload);
  expect(err).toBeFalsy();

  const message = UserRegistration.create(payload);
  const buffer = UserRegistration.encode(message).finish();
  expect(buffer).toBeTruthy();

  const res = await Axios.post("http://127.0.0.1:8081/sign_up", buffer, {
    responseType: "arraybuffer",
  });

  const Response = root.lookupType("Response")
  // const ResponseCode = root.lookupEnum("ResponseCode")

  const data = Response.decode(res.data);

  expect(data.toJSON()).toHaveProperty("code", "SUCCESSFUL");
});

