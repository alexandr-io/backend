import logging
import asyncio
import grpc

import metadata_pb2
import metadata_pb2_grpc
import os
import metadata

class Metadata(metadata_pb2_grpc.MetadataServicer):
    async def Metadata(self, request: metadata_pb2.MetadataRequest, context: grpc.aio.ServicerContext) -> metadata_pb2.MetadataReply:
        response = metadata.get_metadata(request.Title, request.Authors)
        return metadata_pb2.MetadataReply(**response)


async def serve() -> None:
    server = grpc.aio.server()
    metadata_pb2_grpc.add_MetadataServicer_to_server(Metadata(), server)
    listen_addr = '[::]:' + os.environ["GRPC_PORT"]
    server.add_insecure_port(listen_addr)
    logging.info("Starting server on %s", listen_addr)
    await server.start()
    try:
        await server.wait_for_termination()
    except Exception:
        await server.stop(0)


if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO)
    asyncio.run(serve())