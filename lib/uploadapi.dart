import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'package:dart_toast/dart_toast.dart';
import 'package:mango_ui/keys.dart';
import 'package:mango_ui/requester.dart';

import 'bodies/up.dart';

String _imageURL;

Future<HttpRequest> createUpload(FormData data) async {
  var apiroute = getEndpoint("artifact");
  var url = "${apiroute}/upload";

  return invokeFormservice(url, data);
}

void uploadFile(Event e) {
  if (e.target is FileUploadInputElement) {
    FileUploadInputElement fileElem = e.target;
    var files = fileElem.files;
    var forAttr = fileElem.dataset['for'];
    var nameAttr = fileElem.dataset['name'];
    var ctrlID = fileElem.id;
    var infoObj = new Up(forAttr, nameAttr, getObjKey());

    if (files.isNotEmpty) {
      File firstFile = files[0];
      doUpload(firstFile, infoObj, ctrlID);
    }
  }
}

void doUpload(File file, Up infoObj, String ctrlID) async {
  var formData = new FormData();
  formData.appendBlob("file", file);
  formData.append("info", jsonEncode(infoObj));

  var req = await createUpload(formData);

  if (req.status == 200) {
    final resp = jsonDecode(req.response);
    finishUpload(resp, infoObj, ctrlID);
  } else {
    new Toast.error(
        title: "Upload Error!",
        message: req.response,
        position: ToastPos.bottomLeft);
  }
}

void finishUpload(dynamic obj, Up infoObj, String ctrlID) async {
  if (_imageURL?.isEmpty ?? true) {
    var apiroute = getEndpoint("artifact");
    _imageURL = "${apiroute}/download";
  }

  var fullURL = "${_imageURL}/${obj}";

  var imageHolder = querySelector("#${ctrlID.replaceFirst('Img', 'View')}");
  var uploader = querySelector("#${ctrlID}");

  imageHolder.classes.remove('is-hidden');
  imageHolder.setAttribute('src', fullURL);

  uploader.dataset['id'] = obj;
  uploader.attributes.remove('required');
}

Future<HttpRequest> removeUpload(String key) async {
  var apiroute = getEndpoint("artifact");
  var url = "${apiroute}/upload/${key}";

  return invokeService("DELETE", url, "");
}
