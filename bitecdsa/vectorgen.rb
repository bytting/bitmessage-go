require 'openssl'
require 'securerandom'
require 'base64'

def vector(failing=false)
    ec = OpenSSL::PKey::EC.new("secp256k1")
    key = ec.generate_key
    digest = OpenSSL::Digest::SHA256.digest(SecureRandom.random_bytes(64))
    sig = OpenSSL::ASN1.decode(key.dsa_sign_asn1(digest))

    keyhex = key.public_key.to_bn.to_s(16).downcase
    keylen = keyhex.length - 2

    x = keyhex[2, keylen/2]
    y = keyhex[2+keylen/2, keylen/2]
    r = sig.value[0].value
    s = sig.value[1].value

    if failing
        r += 1
        s -= 1
    end

    "{\n\t%s,\n\t%s,\n\t%s,\n\t%s,\n\t%s,\n\t%s,\n}," % [Base64.strict_encode64(digest).inspect, x.inspect, y.inspect, r.to_s(16).downcase.inspect, s.to_s(16).downcase.inspect, !failing]
end

10.times do
    puts vector
end

10.times do
    puts vector(true)
end
