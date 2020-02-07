#coding=utf-8
from flask import Blueprint, request, json, jsonify
from model.type import FailResp,SuccessResp, createArchive
from sql.insert import insert

write = Blueprint('write', __name__)

@write.route('/api/write', methods=['POST'])
def index():
    insert(createArchive(request.json))
    return jsonify(SuccessResp())